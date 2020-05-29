package redis

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"order-matching/internal/config"
)

const (
	// redis configuration dependency constants
	_env            = "test"
	_configFilePath = "../../../"

	// test constants
	_testKey1 = "test_key1"
	_testVal1 = "test_value1"
	_testKey2 = "test_key2"
	_testVal2 = "test_value2"
	_testKey3 = "test_key3"
	_testVal3 = `{
    address: '{{integer(100, 999)}} {{street()}}, {{city()}}, {{state()}}, {{integer(100, 10000)}}',
    about: '{{lorem(1, "paragraphs")}}',
    registered: '{{date(new Date(2014, 0, 1), new Date(), "YYYY-MM-ddThh:mm:ss Z")}}',
    latitude: '{{floating(-90.000001, 90)}}',
    longitude: '{{floating(-180.000001, 180)}}',
    tags: [
      '{{repeat(7)}}',
      '{{lorem(1, "words")}}'
    ],
    friends: [
      '{{repeat(3)}}',
      {
        id: '{{index()}}',
        name: '{{firstName()}} {{surname()}}'
      }
    ]
  }`
	_testKey4 = "test_key4"
	_testVal4 = ""
)

var (
	_testRedisCfg *config.RedisConfig
	_rdb          *RedisDb
)

// Runs test setup and teardown functions before and after all
// defined test functions. Then runs all defined test functions.
func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()
	afterAll()
	os.Exit(code)
}

func beforeAll() {
	testCfg, _ := config.NewConfig(_configFilePath, _env)
	_testRedisCfg = &(*testCfg).Redis
	_rdb, _ = NewRedisDb(_testRedisCfg)
  _rdb.Clear()
}

func afterAll() {
  _rdb.Clear()
}

// Tests function that verifies that a connection is established.
func TestVerifyConn(t *testing.T) {
	err := _rdb.verifyConn()
	assert.Equal(t, nil, err)
}

// Tests both insert and select functions.
func TestAddGet(t *testing.T) {
	// function under test
	if err := _rdb.Add(_testKey1, _testVal1); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	val1, err := _rdb.Get(_testKey1)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal1, val1, "should insert a key with a small value")

	// function under test
	if err := _rdb.Add(_testKey2, _testVal2); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	val2, err := _rdb.Get(_testKey2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal2, val2, "should insert a key with a small value")

	// function under test
	if err := _rdb.Add(_testKey3, _testVal3); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	val3, err := _rdb.Get(_testKey3)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal3, val3, "should insert a key with a large value")

	// function under test
	if err := _rdb.Add(_testKey4, _testVal4); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	val4, err := _rdb.Get(_testKey4)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal4, val4, "should insert a key with an empty string value")
}

func TestDelete(t *testing.T) {
  // new test
  // setup
  if err := _rdb.Add(_testKey1, _testVal1); err != nil {
		log.Fatal(err.Error())
	}
  // setup assertion
  exists, err := _rdb.Exists(_testKey1)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")
  // function under test
  err = _rdb.Delete(_testKey1)
  if err != nil {
    log.Fatalf(err.Error())
  }
  // final assertion
  exists, err = _rdb.Exists(_testKey1)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.False(t, exists, "should assert deleted key value pair does not exist")

  // new test
  // setup
  if err := _rdb.Add(_testKey2, _testVal2); err != nil {
    log.Fatal(err.Error())
  }
  // setup assertion
  exists, err = _rdb.Exists(_testKey2)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")
  // function under test
  err = _rdb.Delete(_testKey2)
  if err != nil {
    log.Fatalf(err.Error())
  }
  // final assertion
  exists, err = _rdb.Exists(_testKey2)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.False(t, exists, "should assert deleted key value pair does not exist")

  // new test
  // setup
  if err := _rdb.Add(_testKey3, _testVal3); err != nil {
    log.Fatal(err.Error())
  }
  // setup assertion
  exists, err = _rdb.Exists(_testKey3)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")
  // function under test
  err = _rdb.Delete(_testKey3)
  if err != nil {
    log.Fatalf(err.Error())
  }
  // final assertion
  exists, err = _rdb.Exists(_testKey3)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.False(t, exists, "should assert deleted key value pair does not exist")

  // new test
  // setup
  if err := _rdb.Add(_testKey4, _testVal4); err != nil {
    log.Fatal(err.Error())
  }
  // setup assertion
  exists, err = _rdb.Exists(_testKey4)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")
  // function under test
  err = _rdb.Delete(_testKey4)
  if err != nil {
    log.Fatalf(err.Error())
  }
  // final assertion
  exists, err = _rdb.Exists(_testKey4)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.False(t, exists, "should assert deleted key value pair does not exist")
}

func TestExists(t *testing.T) {
  // new test
  // function under test
  exists, err := _rdb.Exists("non_existent_key")
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.False(t, exists, "should assert key does not exist")

  // new test
  // setup
  if err := _rdb.Add(_testKey1, _testVal1); err != nil {
    log.Fatal(err.Error())
  }
  // function under test
  exists, err = _rdb.Exists(_testKey1)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")

  // new test
  // setup
  if err := _rdb.Add(_testKey2, _testVal2); err != nil {
    log.Fatal(err.Error())
  }
  // function under test
  exists, err = _rdb.Exists(_testKey2)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")

  // new test
  // setup
  if err := _rdb.Add(_testKey3, _testVal3); err != nil {
    log.Fatal(err.Error())
  }
  // function under test
  exists, err = _rdb.Exists(_testKey3)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")

  // new test
  // setup
  if err := _rdb.Add(_testKey4, _testVal4); err != nil {
    log.Fatal(err.Error())
  }
  // function under test
  exists, err = _rdb.Exists(_testKey4)
  if err != nil {
    log.Fatalf(err.Error())
  }
  assert.True(t, exists, "inserted key value pair should exist")
}
