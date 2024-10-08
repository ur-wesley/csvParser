package main

import (
	"errors"
	"os"

	"github.com/sqweek/dialog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Columns      []Column `yaml:"columns"`
	Output       string   `yaml:"output,omitempty"`
	Delimiter    string   `yaml:"delimiter,omitempty"`
	IgnoreHeader bool     `yaml:"ignore_header,omitempty"`
}

type Column struct {
	Column  string            `yaml:"column,omitempty"`
	Name    string            `yaml:"name,omitempty"`
	Index   int               `yaml:"index,omitempty"`
	Suffix  string            `yaml:"suffix,omitempty"`
	Prefix  string            `yaml:"prefix,omitempty"`
	Replace map[string]string `yaml:"replace,omitempty"`
}

func GetConfig() (Config, error) {
	files, err := findConfigFiles()
	if err != nil {
		return Config{}, err
	}

	filename, err := selectConfigFile(files)
	if err != nil {
		return Config{}, err
	}

	configFile, err := readConfigFile(filename)
	if err != nil {
		return Config{}, err
	}

	config, err := parseConfigFile(configFile, filename)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func findConfigFiles() ([]string, error) {
	files := []string{}
	dir, err := os.ReadDir(".")
	if err != nil {
		return nil, errors.New("failed to read directory: " + err.Error())
	}
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		if len(file.Name()) >= 10 && file.Name()[:6] == "config" && file.Name()[len(file.Name())-4:] == ".yml" {
			files = append(files, file.Name())
		}
	}
	if len(files) == 0 {
		dialog.Message("Keine config Datei gefunden").Error()
		return nil, errors.New("no config file found")
	}
	return files, nil
}

func selectConfigFile(files []string) (string, error) {
	var filename string
	if len(files) == 1 {
		filename = files[0]
	} else {
		dialog.Message("Mehrere config Datein gefunden, wähle die richtige aus").Info()
		var err error
		filename, err = dialog.File().Filter("config files", "yml").Title("Konfigurationsdatei auswählen").Load()
		if err != nil {
			return "", errors.New("failed to load config file: " + err.Error())
		}
	}
	return filename, nil
}

func readConfigFile(filename string) ([]byte, error) {
	configFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("failed to read config file " + filename + ": " + err.Error())
	}
	return configFile, nil
}

func parseConfigFile(configFile []byte, filename string) (Config, error) {
	var config Config
	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, errors.New("failed to unmarshal config file " + filename + ": " + err.Error())
	}
	return config, nil
}
