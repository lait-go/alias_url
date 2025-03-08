package config

import (
	"fmt"
	"log"
	"os"
	erration "retsAPI/serv/error"
	"retsAPI/serv/storage"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"prod"`
	StoragePath string `yaml:"storage_path"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
	Timeout	time.Duration `yaml:"timeout"`
	Idle_timeout time.Duration `yaml:"idle_timeout"`
}

const (
	configPath = "../config/config.yaml"
)

func NewConfig() *Config {
	var cfg Config

	env := os.Getenv("CONFIG_PATH")
	
	if env == "" {
		err := ConfigFileWork(configPath, &cfg)
		if err != nil {
			log.Fatal(err)
		}
	}else{
		err := ConfigFileWork(env, &cfg)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &cfg
}

func ConfigFileWork(env string, cfg *Config) error{

	if err := storage.FileExists(env); err == nil {
		if err = cleanenv.ReadConfig(env, cfg); err != nil {
			error := fmt.Sprintf("ERROR_PATH_READING: %s", env)
			erration.LogError(err, error)
			return err
		}

	}else{
		return err
	}

	return nil
}
