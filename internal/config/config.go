package config

import (
	"flag"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WsServer WsServerConfig `yaml:"websocket_server"`
}

type WsServerConfig struct {
	ReadDeadline int    `yaml:"read_deadline"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	MsgSizeLimit int    `yaml:"message_size_limit"`
	Addr         string `yaml:"address"`
	Path         string `yaml:"path"`
}

func NewConfig() (*Config, error) {
	env, err := parseEnvFlag()
	if err != nil {
		return nil, err
	}

	yamlData, err := loadConfigFile("../../", fileName(env))
	if err != nil {
		return nil, err
	}

	cfg, err := unmarshalYAML(yamlData)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Handles reading the environment command line flag passed in.
// Will return an error if an invalid flag is passed in
// (flag must be "prod", "dev", or "test").
func parseEnvFlag() (string, error) {
	var env string
	flag.StringVar(&env, "GO_ENV", "", "Runtime environment")
	flag.Parse()

	if !validEnvFlag(env) {
		return "", errors.Errorf("Invalid environment flag: %s", env)
	}
	return env, nil
}

func validEnvFlag(env string) bool {
	return env == "prod" || env == "dev" || env == "test"
}

func fileName(env string) string {
	return "config." + env + ".yaml"
}

func loadConfigFile(filePath string, fileName string) ([]byte, error) {
	yamlData, err := ioutil.ReadFile(filePath + fileName)
	if err != nil {
		return nil, errors.Errorf("Error opening file: %s", fileName)
	}
	return yamlData, nil
}

func unmarshalYAML(yamlData []byte) (*Config, error) {
	cfg := &Config{}
	if err := yaml.Unmarshal(yamlData, &cfg); err != nil {
		return nil, errors.Errorf("Error unmarshalling yaml")
	}
	return cfg, nil
}
