package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	HttpServer `json:"http_server"`
	GrpcServer `json:"grpc_server"`
}

type HttpServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type GrpcServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var once sync.Once
var instance *Config

func GetConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		file, err := os.Open(path)
		if err != nil {
			log.Fatal("unable to open config file.", err)
		}
		defer file.Close()

		cfg := json.NewDecoder(file)
		err = cfg.Decode(&instance)
		if err != nil {
			log.Fatal("unable to parse config file.", err)
		}
	})
	return instance
}
