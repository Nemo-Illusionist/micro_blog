package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	ConnectionString  string        `json:"connection_string"`
	Address           string        `json:"address"`
	PageSize          int           `json:"page_size"`
	PageSizeMax       int           `json:"page_size_max"`
	PasswordSecretKey string        `json:"password_secret_key"`
	TokenSecretKey    string        `json:"token_secret_key"`
	TokenExpHour      time.Duration `json:"token_exp_hour"`
}

var config *Config = nil

func GetConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	configPath := os.Getenv("CONFIG_PATH")

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return nil, err
	}

	cnf := &Config{}
	err = json.Unmarshal(bs, cnf)
	if err != nil {
		return nil, err
	}

	config = cnf
	return config, nil
}
