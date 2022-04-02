package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJSONConfig retrieves a JSON configuration.
func ReadJSONConfig() (*Config, error) {
	homePath, ok := os.LookupEnv("HOME")
	if !ok {
		return nil, fmt.Errorf("unknown home path")
	}

	configPath := fmt.Sprintf("%s/.config/cappuccino.json", homePath)
	_, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
