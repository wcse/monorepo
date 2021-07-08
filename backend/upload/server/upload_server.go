package server

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

//
// api /upload
//
func (s *server) upload(w http.ResponseWriter, r *http.Request) {
	s.logger.Info(fmt.Sprintf("upload: Method: %s, Length: %d", r.Method, r.ContentLength))

	if r.Method != http.MethodPost {
		http.Error(w, "not supported method "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get file, handler from key file
	file, handler, err := r.FormFile(fileFormKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	contentType := handler.Header.Get(contentTypeHeader)

	fileType := supportFileTypes[contentType]
	if fileType == "" || fileType == NA {
		http.Error(
			w,
			fmt.Sprintf("do not support with type %s", contentType),
			http.StatusBadRequest,
		)
		return
	}

	var response *Response
	switch fileType {
	case ZIP:
		response = unzipAndUploadFiles(s.logger, file, handler)
	case IMAGE:
		response = uploadFile(s.logger, file, handler)
	default:
		response = &Response{
			Error: fmt.Sprintf("unhandle for file type %s", contentType),
		}
	}

	// make json
	result, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set content-type and return
	w.Header().Set(contentTypeHeader, jsonType)
	_, err = w.Write(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//
// uploadFile
//
func uploadFile(
	logger *zap.Logger,
	file multipart.File,
	handler *multipart.FileHeader) *Response {

	suffix := createSuffix()

	ext := filepath.Ext(handler.Filename)
	contentType := mime.TypeByExtension(ext) // get file type by extension
	fileName := fmt.Sprintf(
		"%s_%s%s",
		handler.Filename[0:len(handler.Filename)-len(ext)],
		suffix,
		ext,
	)

	logger.Info(fmt.Sprintf("uploadFile: FileName: %s, Type: %s", fileName, contentType))

	buffer := make([]byte, handler.Size)
	_, err := file.Read(buffer)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	_, err = uploadFileToCloud(fileName, &buffer)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	// todo make struct common response
	return &Response{
		Success: true,
		Data: map[string]interface{}{
			"baseDirectoryName": "",
			"result": map[string]interface{}{
				"fileName": fileName,
				"size":     len(buffer),
				"type":     contentType,
				"path":     fileName,
			},
		},
	}
}

//
// unzipAndUploadFiles
// unzip and upload all files inside
// note*: only upload the first folder (if exist) after extract
func unzipAndUploadFiles(
	logger *zap.Logger,
	file multipart.File,
	handler *multipart.FileHeader) *Response {

	suffix := createSuffix()

	// unzip and push to s3
	unzipped, err := zip.NewReader(file, handler.Size)
	if err != nil {
		return &Response{
			Error: err.Error(),
		}
	}

	var errors []map[string]interface{}
	var results []map[string]interface{}

	var baseDir string

	for _, f := range unzipped.File {
		dirNames := strings.Split(f.Name, string(os.PathSeparator))

		if f.FileInfo().IsDir() {
			// now we just support only one base directory,
			// it mean when unzip the file, existed only one folder inside
			if baseDir == "" && len(dirNames) == 2 { // => base dir, values of dirNames will be ["$name", ""]
				baseDir = dirNames[0]
			}
			continue
		}

		rc, _ := f.Open()

		content, _ := ioutil.ReadAll(rc)
		name := f.FileInfo().Name()
		ext := filepath.Ext(name)
		contentType := mime.TypeByExtension(ext) // get file type by extension

		// some cases we cannot get mime type from extension (ex: .atlas)
		// => using http.DetectContentType
		if contentType == "" {
			contentType = http.DetectContentType(content)
		}

		fileType := supportFileTypes[contentType]
		if fileType == "" || fileType == NA {
			rc.Close()
			continue
		}

		if fileType == ZIP {
			errors = append(errors, map[string]interface{}{
				"error":    "do not support multiple layer compression",
				"fileName": name,
				"path":     f.Name,
				"type":     contentType,
			})
			rc.Close()
			continue
		}

		// create fileName
		if len(dirNames) == 1 { // only file name, no dir
			dirNames[0] = fmt.Sprintf(
				"%s_%s%s",
				name[0:len(name)-len(ext)],
				suffix,
				ext,
			)
		} else { //
			// some cases file after unzip not return directory first
			// todo: try other way to get root dir when extract file
			if baseDir == "" {
				baseDir = dirNames[0]

			} else if baseDir != dirNames[0] {
				logger.Debug(fmt.Sprintf("can not get base directory for file %s", f.Name))
				rc.Close()
				continue
			}

			dirNames[0] = fmt.Sprintf("%s_%s", baseDir, suffix)
		}

		fileUploadName := strings.Join(dirNames, string(os.PathSeparator))
		logger.Info(fmt.Sprintf("uploadFile: FileName: %s", fileUploadName))

		_, err = uploadFileToCloud(fileUploadName, &content)
		if err != nil {
			errors = append(errors, map[string]interface{}{
				"error":    fmt.Sprintf("Upload error: %s", err.Error()),
				"fileName": name, // return real name and path for easy checking file upload failed
				"path":     f.Name,
				"type":     contentType,
			})
		} else {
			results = append(results, map[string]interface{}{
				"fileName": dirNames[len(dirNames)-1],
				"type":     contentType,
				"size":     len(content),
				"path":     fileUploadName,
			})
		}
		rc.Close()
	}

	// set baseDir = baseDir + suffix
	if baseDir != "" {
		baseDir = fmt.Sprintf("%s_%s", baseDir, suffix)
	}

	logger.Info(fmt.Sprintf("paths %s", baseDir))
	success := len(results) > 0
	// todo make struct common response
	return &Response{
		Success: success,
		Data: map[string]interface{}{
			"baseDirectoryName": baseDir,
			"results":           results,
		},
		Error: errors,
	}
}

//
// createSuffix
//
func createSuffix() string {
	return fmt.Sprintf(
		"%s_%s",
		getRandomString(stringSuffixLength),
		time.Now().Format(timeSuffixFormat),
	)
}

// GetString returns a random string
func getRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune(lettersRandom)
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func uploadFileToCloud(name string, content *[]byte) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}
