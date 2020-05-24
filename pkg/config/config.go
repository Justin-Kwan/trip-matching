package config

import (
	"flag"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	_cmdFlag        = "GO_ENV"
	_cmdFlagDefault = "Invalid default environment flag"
	_cmdFlagDesc    = "Application environment"
	_envFilePath    = "../../"
	_fileType       = "yaml"
)

type Config struct {
	WsConfig WsConfig `mapstructure:"websocket_server"`
}

type WsConfig struct {
	ReadDeadline int    `mapstructure:"read_deadline"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	MsgSizeLimit int    `mapstructure:"message_size_limit"`
	Addr         string `mapstructure:"address"`
	Path         string `mapstructure:"path"`
}

func validEnvFlag(env string) bool {
	return env == "prod" || env == "dev" || env == "test"
}

func parseEnvFlag() (string, error) {
	var env string
	flag.StringVar(&env, _cmdFlag, _cmdFlagDefault, _cmdFlagDesc)
	flag.Parse()
	if !validEnvFlag(env) {
		return "", errors.Errorf("Invalid env flag given: %s", env)
	}

	return env, nil
}

func loadConfigFile(env string) error {
	viper.SetConfigType(_fileType)
	viper.SetConfigName("env." + env)
	viper.AddConfigPath(_envFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return errors.Errorf("Error opening file %w", err)
	}
	return nil
}

func unmarshalConfigFile() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Errorf("Error unmarshalling file %w", err)
	}

	return cfg, nil
}

func NewConfig() (*Config, error) {
	env, err := parseEnvFlag()
	if err != nil {
		return nil, err
	}

	if err := loadConfigFile(env); err != nil {
		return nil, err
	}

	cfg, err := unmarshalConfigFile()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
