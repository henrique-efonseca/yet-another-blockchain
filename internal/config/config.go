package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Network struct {
		Name  string   `yaml:"name"`
		Port  int      `yaml:"port"`
		Peers []string `yaml:"peers"`
	} `yaml:"network"`
	Consensus struct {
		Type       string `yaml:"type"`
		Difficulty int    `yaml:"difficulty"`
	} `yaml:"consensus"`
	Block struct {
		MaxSize int `yaml:"maxSize"`
	} `yaml:"block"`
	SmartContracts struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"smartContracts"`
	Wallet struct {
		Enabled bool `yaml:"enabled"`
		Balance int  `yaml:"balance"`
	} `yaml:"wallet"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
