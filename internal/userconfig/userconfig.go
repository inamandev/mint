package userconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type UserConfig struct {
	GitHub   GitHub   `json:"github"`
	Defaults Defaults `json:"defaults"`
}

type GitHub struct {
	Username string `json:"username"`
}

type Defaults struct {
	Language string          `json:"language"`
	Arch     string          `json:"arch"`
	DI       string          `json:"di"`
	Features DefaultFeatures `json:"features"`
}

type DefaultFeatures struct {
	Makefile     bool `json:"makefile"`
	DevContainer bool `json:"devcontainer"`
	CI           bool `json:"ci"`
	Dockerfile   bool `json:"dockerfile"`
	Linter       bool `json:"linter"`
	Air          bool `json:"air"`
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	return filepath.Join(home, ".mint", "config.json"), nil
}

// Load reads ~/.mint/config.json. Returns empty config if file does not exist.
func Load() (UserConfig, error) {
	path, err := configPath()
	if err != nil {
		return UserConfig{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return UserConfig{}, nil
		}
		return UserConfig{}, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg UserConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return UserConfig{}, fmt.Errorf("invalid config JSON: %w", err)
	}
	return cfg, nil
}

// Save writes config to ~/.mint/config.json, creating the directory if needed.
func Save(cfg UserConfig) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}
	return nil
}
