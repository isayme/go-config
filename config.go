package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

var once sync.Once

// Parse parse config
func Parse(v interface{}, path ...string) {
	once.Do(func() {
		filePath, err := getConfigFilePath(path...)
		if err != nil {
			panic(err)
		}

		ext := filepath.Ext(filePath)

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		err = parseConfig(data, ext, v)
		if err != nil {
			panic(err)
		}
	})
}

func getConfigFilePath(path ...string) (string, error) {
	var filePath string

	if len(path) > 0 {
		filePath = path[0]
	}

	// default config file path from env
	if filePath == "" {
		filePath = os.Getenv("CONF_FILE_PATH")
	}

	if filePath == "" {
		return "", fmt.Errorf("config file not specified")
	}

	return filePath, nil
}

type unmarshaler func(data []byte, v interface{}) error

func parseConfig(data []byte, ext string, v interface{}) error {
	ext = strings.TrimLeft(ext, ".")

	var unmarshal unmarshaler

	switch ext {
	case "json":
		unmarshal = json.Unmarshal
	case "yaml", "yml":
		unmarshal = yaml.Unmarshal
	default:
		return fmt.Errorf("not supported config ext: %s", ext)
	}

	return unmarshal(data, v)
}
