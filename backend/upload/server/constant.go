package server

const maxFileSize = 20 << 20 // 20MB

const (
	fileFormKey       = "file"
	contentTypeHeader = "Content-Type"
	jsonType          = "application/json"

	timeSuffixFormat   = "2006010215"
	stringSuffixLength = 10
	lettersRandom      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
)

const (
	NA         = "n/a"
	IMAGE      = "image"
	ZIP        = "zip"
	JAVASCRIPT = "javascript"
	JSON       = "json"
	TEXT       = "text"
	CSV        = "csv"
	PDF        = "pdf"
)

// map: content-type support for upload and check need to be unzip or not
// ref: https://github.com/gabriel-vasile/mimetype/blob/master/supported_mimes.md
var supportFileTypes = map[string]string{
	"application/octet-stream":     NA,
	"application/zip":              ZIP,
	"application/x-zip":            ZIP,
	"application/x-zip-compressed": ZIP,
	"image/jpeg":                   IMAGE,
	"image/png":                    IMAGE,
	"image/gif":                    IMAGE,
	"application/javascript":       JAVASCRIPT,
	"application/json":             JSON,
	"text/plain; charset=utf-8":    TEXT,
	"text/csv":                     CSV,
	"application/pdf":              PDF,
}
