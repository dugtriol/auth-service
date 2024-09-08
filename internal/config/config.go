package config

import (
	`os`
	`time`

	`github.com/ilyakaznacheev/cleanenv`
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config doesn't exist: " + path)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("panic to read config: " + err.Error())
	}

	return &cfg
}
