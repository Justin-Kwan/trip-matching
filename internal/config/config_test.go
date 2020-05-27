package config

import (
	// "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidEnvFlag(t *testing.T) {
	assert.Equal(t, true, validEnvFlag("prod"), "should be a valid environment flag")
	assert.Equal(t, true, validEnvFlag("dev"), "should be a valid environment flag")
	assert.Equal(t, true, validEnvFlag("test"), "should be a valid environment flag")
	assert.Equal(t, false, validEnvFlag("prod "), "should be a invalid environment flag")
	assert.Equal(t, false, validEnvFlag(" test"), "should be a invalid environment flag")
	assert.Equal(t, false, validEnvFlag(" dev"), "should be a invalid environment flag")
	assert.Equal(t, false, validEnvFlag("invalid"), "should be a invalid environment flag")
}

func TestFileName(t *testing.T) {
	assert.Equal(t, "config.prod.yaml", fileName("prod"), "should return production configuration file name")
	assert.Equal(t, "config.dev.yaml", fileName("dev"), "should return development configuration file name")
	assert.Equal(t, "config.test.yaml", fileName("test"), "should return test configuration file name")
	assert.Equal(t, "config. .yaml", fileName(" "), "should return configuration file with ' ' in name")
	assert.Equal(t, "config.*.yaml", fileName("*"), "should return configuration file with '*' in name")
}

// func TestUnmarshalYAML(t *testing.T) {
	// YAMLTextData1 := `websocket_server:
	// read_deadline: 1
	// read_timeout: 1
	// write_timeout: 1
	// message_size_limit: 1
	// address: "127.0.0.1:8001"
	// path: "/path1"
	// redis_store:
	// item_expiry:1
	// idle_timeout:1
	// max_idle_connections:1
	// max_active_connections:1
	// address:"127.0.0.1:6001"
	// password:"test_redis_password1"
	// connection_protocol:"tcp1"`

	// YAMLTextData2 := `
  // websocket_server:
  //   read_deadline: 					2
  //   read_timeout: 					2
  //   write_timeout: 					2
  //   message_size_limit: 		2
  //   address: 								"127.0.0.1:8002"
  //   path: 									"/path2
	// redis_store:
	// 	item_expiry:            2
	// 	idle_timeout:           2
	// 	max_idle_connections:   2
	// 	max_active_connections: 2
	// 	address:                "127.0.0.1:6002"
	// 	password:               "test_redis_password2"
	// 	connection_protocol:    "tcp2"`
	//
	// badYAMLTextData3 := `
  // websocket_server:
  //   read_deadline: 					3
  //   read_timeout: 					3
  //   write_timeout: 					"string shouldn't be here"
  //   message_size_limit: 		3
  //   address: 								"127.0.0.1:8003"
  //   path: 									"/path3
	// redis_store:
	// 	item_expiry:            3
	// 	idle_timeout:           3
	// 	max_idle_connections:   3
	// 	max_active_connections: 3
	// 	address:                "127.0.0.1:6003"
	// 	password:               "test_redis_password3"
	// 	connection_protocol:    "tcp3"`
	//
	// badYAMLTextData4 := `
  // websocket_server:
  //   read_deadline: 					"string shouldn't be here"
  //   read_timeout: 					4
  //   write_timeout: 					4
  //   message_size_limit: 		4
  //   address: 								"127.0.0.1:8004"
  //   path: 									"/path4
	// redis_store:
	// 	item_expiry:            4
	// 	idle_timeout:           4
	// 	max_idle_connections:   4
	// 	max_active_connections: 4
	// 	address:                "127.0.0.1:6004"
	// 	password:               "test_redis_password4"
	// 	connection_protocol:    "tcp4"`

	// setup
	// YAMLByteData1 := []byte(YAMLTextData1)
	// YAMLByteData2 := []byte(YAMLTextData2)
	// badYAMLByteData3 := []byte(badYAMLTextData3)
	// badYAMLByteData4 := []byte(badYAMLTextData4)

	// // function under test
	// cfg1, err1 := unmarshalYAML(YAMLByteData1)
	// if err1 != nil {
	// 	t.Error("Error parsing 1st YAML byte data")
	// 	log.Fatalf("Error yaml 1: %v", err1);
	// }
	//
	// assert.Equal(t,
	// 	Config{
	// 		WsServerConfig{
	// 			ReadDeadline: 1,
	// 			ReadTimeout:  1,
	// 			WriteTimeout: 1,
	// 			MsgSizeLimit: 1,
	// 			Addr:         "127.0.0.1:8001",
	// 			Path:         "/path1",
	// 		},
	// 		RedisConfig{
	// 			Exp:          3,
	// 			IdleTimeout:  200,
	// 			MaxIdle:      500,
	// 			MaxActive:    1200,
	// 			Addr:         "127.0.0.1:6379",
	// 			Password:     "test_redis_password",
	// 			ConnProtocol: "tcp",
	// 		},
	// 	}, *cfg1, "should assert config structs equal")

	// function under test
// 	cfg2, err2 := unmarshalYAML(YAMLByteData2)
// 	if err2 != nil {
// 		t.Error("Error parsing 2nd YAML byte data")
// 	}
//
// 	assert.Equal(t,
// 		Config{
// 			WsServerConfig{
// 				ReadDeadline: 2,
// 				ReadTimeout:  2,
// 				WriteTimeout: 2,
// 				MsgSizeLimit: 2,
// 				Addr:         "127.0.0.1:8002",
// 				Path:         "/path2",
// 			},
// 			RedisConfig{
// 				Exp:          2,
// 				IdleTimeout:  2,
// 				MaxIdle:      2,
// 				MaxActive:    2,
// 				Addr:         "127.0.0.1:6002",
// 				Password:     "test_redis_password2",
// 				ConnProtocol: "tcp2",
// 			},
// 		}, *cfg2, "should assert config structs equal")
//
// 	// function under test
// 	_, err3 := unmarshalYAML(badYAMLByteData3)
// 	assert.EqualError(t, err3, "Error unmarshalling yaml")
//
// 	// function under test
// 	_, err4 := unmarshalYAML(badYAMLByteData4)
// 	assert.EqualError(t, err4, "Error unmarshalling yaml")
// }
