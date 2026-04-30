package repl

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	LastLanguage int32 `json:"last_language"`
}

func (state *AppState) GetConfig() (Config, error) {
	//if the file doesnt exist, return empty config with no err
	file, err := os.Open("conf.json")
	if err != nil {
		if os.IsNotExist(err) {
			config := Config{}
			return config, nil
		} else {
			return Config{}, err
		}
	}

	//decode the config
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, err
}

func (state *AppState) SaveConfig(config Config) error {
	file, err := os.OpenFile("conf.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}
