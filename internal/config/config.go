package config

import "github.com/spf13/viper"

type Config struct {
	Server    ServerListen `mapstructure:"server_listen"`
	Tarantool Tarantool    `mapstructure:"tarantool_settings"`
	Tokens    Tokens       `mapstructure:"tokens"`
}

type ServerListen struct {
	IP   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
}

type Tarantool struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	User           string `mapstructure:"user"`
	TimeoutSeconds int    `mapstructure:"timeout_seconds"`
}

type Tokens struct {
	Key string `mapstructure:"key"`
}

func New(path string) (*Config, error) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := &Config{}

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
