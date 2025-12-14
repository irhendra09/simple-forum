package configs

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}

type Config struct {
	JWTSecret string   `yaml:"jwt_secret"`
	Database  DBConfig `yaml:"database"`
}

// AppConfig holds loaded configuration from YAML file. Environment variables
// override values in this struct when getter helpers are called.
var AppConfig *Config

// LoadConfig reads a YAML config file from path (e.g. "config.yaml"). If path
// is empty it will attempt to read "config.yaml" in the working directory.
// If the file doesn't exist, LoadConfig returns nil and AppConfig remains nil.
func LoadConfig(path string) error {
	if path == "" {
		path = "config.yaml"
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		// file not found is not an error for loading optional config
		return nil
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return err
	}
	AppConfig = &cfg
	return nil
}

// GetJWTSecret returns JWT secret – environment variable `JWT_SECRET` has the
// highest priority, then value from AppConfig (if loaded), then fallback.
func GetJWTSecret() string {
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return s
	}
	if AppConfig != nil && AppConfig.JWTSecret != "" {
		return AppConfig.JWTSecret
	}
	return "secret"
}
