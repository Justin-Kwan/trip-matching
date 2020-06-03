package config

import (
	"flag"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WsServer   WsServerConfig `yaml:"websocket_server"`
	RedisKeyDB RedisConfig    `yaml:"redis_keydb"`
	RedisGeoDB RedisConfig    `yaml:"redis_geodb"`
}

type WsServerConfig struct {
	ReadDeadline int    `yaml:"read_deadline"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	MsgSizeLimit int    `yaml:"message_size_limit"`
	Addr         string `yaml:"address"`
	Path         string `yaml:"path"`
}

type RedisConfig struct {
	IdleTimeout  int    `yaml:"idle_timeout"`
	MaxIdle      int    `yaml:"max_idle_connections"`
	MaxActive    int    `yaml:"max_active_connections"`
	Addr         string `yaml:"address"`
	Password     string `yaml:"password"`
	ConnProtocol string `yaml:"connection_protocol"`
}

// Handles reading the environment command line flag passed in.
// Will return an error if an invalid flag is passed in
// (flag must be "prod", "dev", or "test").
func ParseEnvFlag() (string, error) {
	var env string
	flag.StringVar(&env, "GO_ENV", "", "Runtime environment")
	flag.Parse()

	if !validEnvFlag(env) {
		return "", errors.Errorf("Invalid environment flag: %s", env)
	}
	return env, nil
}

// Checks if a string passed in is a correct environment command
// line flag.
func validEnvFlag(env string) bool {
	return env == "prod" || env == "dev" || env == "test"
}

// Returns a new configuration struct containing values needed for
// the websocket and rest server. The specific yaml config file
// parsed into the struct depends on the runtime environment flag
// passed in the command line. The file path is relative to the caller.
func NewConfig(filePath, env string) (*Config, error) {
	yamlData, err := loadConfigFile(filePath, fileName(env))
	if err != nil {
		return nil, err
	}

	cfg, err := unmarshalYAML(yamlData)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Returns the config file name based on the environment flag
// passed in.
func fileName(env string) string {
	return "config." + env + ".yaml"
}

// Loads and parses a yaml file given the file path and file name.
// Then returns the parsed yaml data in a byte array.
func loadConfigFile(filePath, fileName string) ([]byte, error) {
	yamlData, err := ioutil.ReadFile(filePath + fileName)
	if err != nil {
		return nil, errors.Errorf("Error opening file '%s': %v", fileName, err)
	}
	return yamlData, nil
}

// Converts config yaml data (in a byte array) into a Config struct.
// Then returns the config struct.
func unmarshalYAML(yamlData []byte) (*Config, error) {
	cfg := &Config{}
	if err := yaml.Unmarshal(yamlData, &cfg); err != nil {
		return nil, errors.Errorf("Error unmarshalling yaml: %v", err)
	}
	return cfg, nil
}
