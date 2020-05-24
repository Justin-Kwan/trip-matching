package config

import (
	"flag"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	_cmdFlag        = "GO_ENV"
	_cmdFlagDefault = "Default flag"
	_cmdFlagDesc    = "Application environment"
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

type FileInfo struct {
	Name string
	Type string
	Path string
}

// Handles reading the environment command line flag passed in.
// Will return an error if an invalid flag is passed in
// (flag must be "prod", "dev", or "test").
func parseEnvFlag() (string, error) {
	var env string
	flag.StringVar(&env, _cmdFlag, _cmdFlagDefault, _cmdFlagDesc)
	flag.Parse()
	if !validEnvFlag(env) {
		return "", errors.Errorf("Invalid env flag given: %s", env)
	}
	return env, nil
}

func validEnvFlag(env string) bool {
	return env == "prod" || env == "dev" || env == "test"
}

func newFileInfo(env string, fileType string, filePath string) *FileInfo {
	return &FileInfo{
		Name: "config." + env,
		Type: fileType,
		Path: filePath,
	}
}

func loadConfigFile(fi *FileInfo) error {
	viper.SetConfigName(fi.Name)
	viper.SetConfigType(fi.Type)
	viper.AddConfigPath(fi.Path)

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
	return &cfg, nil
}

func NewConfig() (*Config, error) {
	env, err := parseEnvFlag()
	if err != nil {
		return nil, err
	}

	fi := newFileInfo(env, "yaml", "../../")
	if err := loadConfigFile(fi); err != nil {
		return nil, err
	}

	cfg, err := unmarshalConfigFile()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", *cfg)

	return cfg, nil
}
