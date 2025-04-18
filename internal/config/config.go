package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSL_Mode string `yaml:"sslmode"`
}

type HttpServerConfig struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	Idle_Timeout time.Duration `yaml:"idle_timeout"`
}

type RedisServerConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Status   int    `yaml:"worflow_status"`
}

type Config struct {
	Env         string            `yaml:"env"`
	SQLite_path string            `yaml:"sqlite_path"`
	Postgres    PostgresConfig    `yaml:"postgres"`
	Redis       RedisServerConfig `yaml:"redis"`
	HttpServer  HttpServerConfig  `yaml:"http_server"`
}

func MustLoad() *Config {
	config_path := os.Getenv("CONFIG_PATH")

	if config_path == "" {
		log.Fatal("[ config.go ] Config_path is not set")
	}
	if _, err := os.Stat(config_path); os.IsExist(err) {
		log.Fatalf("[ config.go ] Config is not exist: %s\n", config_path)
	}

	var config Config

	if err := cleanenv.ReadConfig(config_path, &config); err != nil {
		log.Fatalf("[ config.go ] Cannot read config: %s\n", config_path)
	}

	return &config
}

func (p PostgresConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSL_Mode,
	)
}
