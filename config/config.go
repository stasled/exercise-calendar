package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HttpServer `mapstructure:"http_server"`
	GrpcServer `mapstructure:"grpc_server"`
	DB         `mapstructure:"db"`
}

type HttpServer struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	TimeoutRead  time.Duration `mapstructure:"timeout_read"`
	TimeoutWrite time.Duration `mapstructure:"timeout_write"`
	TimeoutIdle  time.Duration `mapstructure:"timeout_idle"`
}

type GrpcServer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DB struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

func GetConfig(path string) *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &cfg
}
