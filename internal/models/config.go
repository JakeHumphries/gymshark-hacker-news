package models

import (
	"errors"
	"os"
)

type Config struct {
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabasePort     string
	Cron             string
	WorkerCount      int
	ApiHost          string
	ApiPort          string
}

func GetConfig() (*Config, error) {
	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return nil, errors.New("err: env var database user doesnt exist")
	}
	pass, exists := os.LookupEnv("DB_PASS")
	if !exists {
		return nil, errors.New("err: env var database pass doesnt exist")
	}
	name, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return nil, errors.New("err: env var database name doesnt exist")
	}
	port, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return nil, errors.New("err: env var database port doesnt exist")
	}
	host, exists := os.LookupEnv("API_HOST")
	if !exists {
		host = "0.0.0.0"
	}

	apiPort, exists := os.LookupEnv("API_PORT")
	if !exists {
		apiPort = "8000"
	}

	config := Config{
		DatabaseName:     name,
		DatabaseUser:     user,
		DatabasePassword: pass,
		DatabasePort:     port,
		Cron:             "0 30 * * * *",
		WorkerCount:      10,
		ApiHost:          host,
		ApiPort:          apiPort,
	}

	return &config, nil
}
