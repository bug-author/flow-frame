package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	LLMProvider  string `json:"llmProvider"`
	Model        string `json:"model"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Port         int    `json:"port"`
}

func LoadConfig(confiFile string) (Config, error) {
	file, err := os.ReadFile(confiFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	return config, err
}
