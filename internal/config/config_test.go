package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

func TestValidEnvFlag(t *testing.T) {
	assert.Equal(t, validEnvFlag("prod"), true, "should be a valid environment flag")
	assert.Equal(t, validEnvFlag("dev"), true, "should be a valid environment flag")
	assert.Equal(t, validEnvFlag("test"), true, "should be a valid environment flag")
	assert.Equal(t, validEnvFlag("prod "), false, "should be a invalid environment flag")
	assert.Equal(t, validEnvFlag(" test"), false, "should be a invalid environment flag")
	assert.Equal(t, validEnvFlag(" dev"), false, "should be a invalid environment flag")
	assert.Equal(t, validEnvFlag("invalid"), false, "should be a invalid environment flag")
}

func TestFileName(t *testing.T) {
	assert.Equal(t, fileName("prod"), "config.prod.yaml", "should return production configuration file name")
	assert.Equal(t, fileName("dev"), "config.dev.yaml", "should return development configuration file name")
	assert.Equal(t, fileName("test"), "config.test.yaml", "should return test configuration file name")
	assert.Equal(t, fileName(" "), "config. .yaml", "should return configuration file with ' ' in name")
	assert.Equal(t, fileName("*"), "config.*.yaml", "should return configuration file with '*' in name")
}

func TestUnmarshalYAML(t *testing.T) {
	YAMLTextData1 := `
  websocket_server:
    read_deadline: 1
    read_timeout: 1
    write_timeout: 1
    message_size_limit: 1
    address: "127.0.0.1:8001"
    path: "/path1"`

	YAMLTextData2 := `
  websocket_server:
    read_deadline: 2
    read_timeout: 2
    write_timeout: 2
    message_size_limit: 2
    address: "127.0.0.1:8002"
    path: "/path2"`

	badYAMLTextData3 := `
  websocket_server:
    read_deadline: 3
    read_timeout: 3
    write_timeout: "string shouldn't be here"
    message_size_limit: 3
    address: "127.0.0.1:8003"
    path: "/path3"`

  badYAMLTextData4 := `
  websocket_server:
    read_deadline: "string shouldn't be here"
    read_timeout: 4
    write_timeout: 4
    message_size_limit: 4
    address: "127.0.0.1:8004"
    path: "/path4"`

  // setup
	YAMLByteData1 := []byte(YAMLTextData1)
	YAMLByteData2 := []byte(YAMLTextData2)
	badYAMLByteData3 := []byte(badYAMLTextData3)
  badYAMLByteData4 := []byte(badYAMLTextData4)

	// function under test
	cfg1, err1 := unmarshalYAML(YAMLByteData1)
	if err1 != nil {
		t.Error("Error parsing 1st YAML byte data")
	}

	assert.Equal(t,
		Config{
			WsServerConfig{
				ReadDeadline: 1,
				ReadTimeout:  1,
				WriteTimeout: 1,
				MsgSizeLimit: 1,
				Addr:         "127.0.0.1:8001",
				Path:         "/path1",
			},
		}, *cfg1, "should assert config structs equal")

	// function under test
	cfg2, err2 := unmarshalYAML(YAMLByteData2)
	if err2 != nil {
		t.Error("Error parsing 2nd YAML byte data")
	}

	assert.Equal(t,
		Config{
			WsServerConfig{
				ReadDeadline: 2,
				ReadTimeout:  2,
				WriteTimeout: 2,
				MsgSizeLimit: 2,
				Addr:         "127.0.0.1:8002",
				Path:         "/path2",
			},
		}, *cfg2, "should assert config structs equal")

	// function under test
	_, err3 := unmarshalYAML(badYAMLByteData3)
	assert.EqualError(t, err3, "Error unmarshalling yaml")

  // function under test
  _, err4 := unmarshalYAML(badYAMLByteData4)
  assert.EqualError(t, err4, "Error unmarshalling yaml")
}
