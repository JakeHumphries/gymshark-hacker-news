package models

import (
	"errors"
	"os"
	"time"
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
	RedisHost        string
	CacheTimout      time.Duration
	GrpcPort         string
	GrpcHost         string
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

	redisHost, exists := os.LookupEnv("REDIS_HOST")
	if !exists {
		return nil, errors.New("err: env var redis host doesnt exist")
	}

	apiHost, exists := os.LookupEnv("API_HOST")
	if !exists {
		apiHost = "0.0.0.0"
	}

	apiPort, exists := os.LookupEnv("API_PORT")
	if !exists {
		apiPort = "8000"
	}

	grpcPort, exists := os.LookupEnv("GRPC_PORT")
	if !exists {
		grpcPort = "9000"
	}

	grpcHost, exists := os.LookupEnv("GRPC_HOST")
	if !exists {
		return nil, errors.New("err: env var grpc host doesnt exist")
	}

	config := Config{
		DatabaseName:     name,
		DatabaseUser:     user,
		DatabasePassword: pass,
		DatabasePort:     port,
		Cron:             "0 30 * * * *",
		WorkerCount:      10,
		ApiHost:          apiHost,
		ApiPort:          apiPort,
		RedisHost:        redisHost,
		CacheTimout:      5 * time.Minute,
		GrpcPort:         grpcPort,
		GrpcHost:         grpcHost,
	}

	return &config, nil
}
