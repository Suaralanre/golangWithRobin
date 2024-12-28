package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type AppConfig struct {
	Host       string `json:"hostname"`
	Port       uint16 `json:"port"`
	AdminEmail string `json:"admin_email"`
}

func ReadConfig(filename string) (*AppConfig, error) {

	if filename == "" {
		filename = GetDefaultConfigLocation()
	}

	var jsonConfig AppConfig
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&jsonConfig)
	if err != nil {
		return nil, err
	}

	return &jsonConfig, nil
}

func DefaultConfig() *AppConfig {
	return &AppConfig{
		Host:       "localhost",
		Port:       2400,
		AdminEmail: "admin@localhost",
	}
}

func GetDefaultConfigLocation() string {
	if runtime.GOOS == "windows" {
		fmt.Println("We're on windows!")
		return "C:\\Local\\jarvishttp_config.json"
	} else if runtime.GOOS == "linux" {
		fmt.Println("We're on linux!")
		return filepath.Join("/opt", "jarvishttp", "jarvishttp_config.json")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("We're on MacOS (darwin)!")
		return filepath.Join("/opt", "jarvishttp_config.json")
	}

	return ""
}
