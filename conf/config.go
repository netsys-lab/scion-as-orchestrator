package conf

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Command string
}

func NewConfig() *Config {
	return &Config{
		Command: "",
	}
}

func LoadConfig(path string) (*Config, error) {
	c := NewConfig()

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
