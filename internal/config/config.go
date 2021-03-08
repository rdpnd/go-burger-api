package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

// Configuration structure taken from https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp

var Logger = log.New(os.Stdout, "Burger-api: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

type Config struct {
	Web struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
		} `yaml:"timeout"`
	} `yaml:"web"`

	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
}

type Flags struct {
	Path    string
	Migrate bool
}

func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ParseFlags() (Flags, error) {
	// String that contains the configured configuration path
	var configPath string
	var migrate *bool

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yaml", "path to config file")
	migrate = flag.Bool("migrate", false, "perform db migration")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := validateConfigPath(configPath); err != nil {
		return Flags{}, err
	}

	// Return the configuration path
	return Flags{
		Path:    configPath,
		Migrate: *migrate,
	}, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' was provided, but file was expected", path)
	}
	return nil
}
