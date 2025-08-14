package config

import (
	"os"

	"gopkg.in/yaml.v3"
)
const configPath = "./pkg/config/data/config.yaml"


type Config struct {
	Server struct {
        Port int `yaml:"port"`
    } `yaml:"server"`
	Database struct { 
		Host     string `yaml:"host"`
        Port     int `yaml:"port"`
        Username string `yaml:"username"`
        Password string `yaml:"password"`
        Dbname string `yaml:"dbName"`
    } `yaml:"database"`
}

func GetConfig() *Config {
	var config Config
    err := LoadConfig(configPath, &config)
    if err != nil {
        panic(err)
    }
    return &config
}

func LoadConfig(path string, config *Config) error {
	file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
	
    decoder := yaml.NewDecoder(file)
    return decoder.Decode(config)
}
