package redis

import (
	"log"
	"os"
	"strconv"
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
	_rdb *RedisDb
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
	testRedisCfg := &(*testCfg).Redis
	_rdb, _ = NewRedisDb(testRedisCfg)
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
func TestInsert(t *testing.T) {
	// function under test
	if err := _rdb.Insert(_testKey1, _testVal1); err != nil {
		log.Fatal(err.Error())
	}
	val1, err := _rdb.Select(_testKey1)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal1, val1, "should insert a key with a small value")

	// function under test
	if err := _rdb.Insert(_testKey2, _testVal2); err != nil {
		log.Fatal(err.Error())
	}
	val2, err := _rdb.Select(_testKey2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal2, val2, "should insert a key with a small value")

	// function under test
	if err := _rdb.Insert(_testKey3, _testVal3); err != nil {
		log.Fatal(err.Error())
	}
	val3, err := _rdb.Select(_testKey3)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal3, val3, "should insert a key with a large value")

	// function under test
	if err := _rdb.Insert(_testKey4, _testVal4); err != nil {
		log.Fatal(err.Error())
	}
	val4, err := _rdb.Select(_testKey4)
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, _testVal4, val4, "should insert a key with an empty string value")

	_rdb.Clear()
}

func TestSelect(t *testing.T) {
	// function under test (selecting non-existent key)
	_, err := _rdb.Select("non_existent_key")
	assert.EqualError(t, err, "Error getting value using key 'non_existent_key': redigo: nil returned")

	// function under test (selecting non-existent key)
	_, err = _rdb.Select("test_key 1")
	assert.EqualError(t, err, "Error getting value using key 'test_key 1': redigo: nil returned")

	// function under test (selecting non-existent key)
	_, err = _rdb.Select(" test_key2")
	assert.EqualError(t, err, "Error getting value using key ' test_key2': redigo: nil returned")

	// function under test (selecting non-existent key)
	_, err = _rdb.Select("test_key3 ")
	assert.EqualError(t, err, "Error getting value using key 'test_key3 ': redigo: nil returned")

	_rdb.Clear()
}

func TestDelete(t *testing.T) {
	// new test
	// setup
	if err := _rdb.Insert(_testKey1, _testVal1); err != nil {
		log.Fatal(err.Error())
	}
	// setup assertion
	assert.True(t, _rdb.Exists(_testKey1), "inserted key value pair should exist")
	// function under test
	err := _rdb.Delete(_testKey1)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// final assertion
	assert.False(t, _rdb.Exists(_testKey1), "should assert deleted key value pair does not exist")

	// new test
	// setup
	if err := _rdb.Insert(_testKey2, _testVal2); err != nil {
		log.Fatal(err.Error())
	}
	// setup assertion
	assert.True(t, _rdb.Exists(_testKey2), "inserted key value pair should exist")
	// function under test
	err = _rdb.Delete(_testKey2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// final assertion
	assert.False(t, _rdb.Exists(_testKey2), "should assert deleted key value pair does not exist")

	// new test
	// setup
	if err := _rdb.Insert(_testKey3, _testVal3); err != nil {
		log.Fatal(err.Error())
	}
	// setup assertion
	assert.True(t, _rdb.Exists(_testKey3), "inserted key value pair should exist")
	// function under test
	err = _rdb.Delete(_testKey3)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// final assertion
	assert.False(t, _rdb.Exists(_testKey3), "should assert deleted key value pair does not exist")

	// new test
	// setup
	if err := _rdb.Insert(_testKey4, _testVal4); err != nil {
		log.Fatal(err.Error())
	}
	// setup assertion
	assert.True(t, _rdb.Exists(_testKey4), "inserted key value pair should exist")
	// function under test
	err = _rdb.Delete(_testKey4)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// final assertion
	assert.False(t, _rdb.Exists(_testKey4), "should assert deleted key value pair does not exist")

	_rdb.Clear()
}

func TestExists(t *testing.T) {
	// new test
	// function under test
	assert.False(t, _rdb.Exists("non_existent_key"), "should assert key does not exist")

	// new test
	// setup
	if err := _rdb.Insert(_testKey1, _testVal1); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	assert.True(t, _rdb.Exists(_testKey1), "inserted key value pair should exist")

	// new test
	// setup
	if err := _rdb.Insert(_testKey2, _testVal2); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	assert.True(t, _rdb.Exists(_testKey2), "inserted key value pair should exist")

	// new test
	// setup
	if err := _rdb.Insert(_testKey3, _testVal3); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	assert.True(t, _rdb.Exists(_testKey3), "inserted key value pair should exist")

	// new test
	// setup
	if err := _rdb.Insert(_testKey4, _testVal4); err != nil {
		log.Fatal(err.Error())
	}
	// function under test
	assert.True(t, _rdb.Exists(_testKey4), "inserted key value pair should exist")

	_rdb.Clear()
}

func insertKeys(keyCount int) {
	for i := 0; i < keyCount; i++ {
		_rdb.Insert(strconv.Itoa(i), "test_value")
	}
}

func TestCountKeys(t *testing.T) {
	_rdb.Clear()

	// new test
	// setup
	insertKeys(1)
	// function under test
	keyCount1, err := _rdb.CountKeys()
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, 1, keyCount1, "key count should be 1")

	_rdb.Clear()

	// new test
	// setup
	insertKeys(2)
	// function under test
	keyCount2, err := _rdb.CountKeys()
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, 2, keyCount2, "key count should be 2")

	_rdb.Clear()

	// new test
	// setup
	insertKeys(3)
	// function under test
	keyCount3, err := _rdb.CountKeys()
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, 3, keyCount3, "key count should be 3")

	_rdb.Clear()

	// new test
	// setup
	insertKeys(4)
	// function under test
	keyCount4, err := _rdb.CountKeys()
	if err != nil {
		log.Fatalf(err.Error())
	}
	assert.Equal(t, 4, keyCount4, "key count should be 4")

	_rdb.Clear()
}
