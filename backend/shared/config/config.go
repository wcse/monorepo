package config

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/golang/protobuf/jsonpb"
	"gopkg.in/yaml.v3"

	"github.com/golang/protobuf/proto"
)

// Flags ...
type Flags struct {
	ConfigPath string
	Template   bool
	Validate   bool
}

// ParseFlags command line arguments.
func ParseFlags() *Flags {
	f := &Flags{}
	flag.StringVar(&f.ConfigPath, "c", "config.yaml", "path to YAML configuration")
	flag.BoolVar(&f.Template, "template", false, "executes go templates on the configuration file")
	flag.BoolVar(&f.Validate, "validate", false, "validates the configuration file and exits")
	flag.Parse()
	return f
}

// ParseFile ...
func ParseFile(path string, pb proto.Message, template bool) error {
	// Get absolute path representation for better error message in case file not found.
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Read file.
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Execute templates if enabled.
	if template {
		contents, err = executeTemplate(contents)
		if err != nil {
			return err
		}
	}

	// Interpolate environment variables.
	contents = []byte(os.ExpandEnv(string(contents)))

	return parseYAML(contents, pb)
}

func parseYAML(contents []byte, pb proto.Message) error {
	// Decode YAML.
	var rawConfig map[string]interface{}
	if err := yaml.Unmarshal(contents, &rawConfig); err != nil {
		return err
	}

	// Encode YAML to JSON.
	jsonBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBuffer).Encode(rawConfig); err != nil {
		return err
	}

	// Unmarshal JSON to proto object.
	if err := jsonpb.Unmarshal(jsonBuffer, pb); err != nil {
		return err
	}

	// All good!
	return nil
}

func executeTemplate(contents []byte) ([]byte, error) {
	tmpl := template.New("config").Funcs(map[string]interface{}{
		"getenv": os.Getenv,
		"getboolenv": func(key string) bool {
			b, _ := strconv.ParseBool(os.Getenv(key))
			return b
		},
	})

	tmpl, err := tmpl.Parse(string(contents))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, nil); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
