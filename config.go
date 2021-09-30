package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

func newConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(filepath.Clean(configPath))
	ExitIfError(err)
	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()
	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	Server struct {
		Verbosity string `yaml:"verbosity"`
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		Timeout   struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
	Application struct {
		Version string `yaml:"version"`
	} `yaml:"application"`
}
