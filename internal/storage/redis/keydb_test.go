package redis

import (
	"log"
	"strconv"
	"testing"

	. "github.com/franela/goblin"
	"github.com/pkg/errors"

	"order-matching/internal/config"
)

const (
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

func insertKeys(keyCount int, keyDB *KeyDB) {
	for i := 0; i < keyCount; i++ {
		keyDB.Insert(strconv.Itoa(i), "test_value")
	}
}

func TestKeyDB(t *testing.T) {
	g := Goblin(t)

	env := "test"
	configFilePath := "../../../"
	dbNum := 0

	var keyDB *KeyDB

	g.Describe("keystore.go tests", func() {

		g.Before(func() {
			testCfg, _ := config.NewConfig(configFilePath, env)
			testRedisCfg := &(*testCfg).Redis
			redisPool, _ := NewPool(testRedisCfg)
			keyDB = NewKeyDB(redisPool, dbNum)
			keyDB.Clear()
		})

		g.AfterEach(func() {
			keyDB.Clear()
		})

		g.Describe("Insert() Tests", func() {
			g.It("should insert a key with a small value", func() {
				// function under test
				if err := keyDB.Insert(_testKey1, _testVal1); err != nil {
					log.Fatal(err.Error())
				}
				val, err := keyDB.Select(_testKey1)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(val).Equal(_testVal1)
			})

			g.It("should insert a key with a small value", func() {
				// function under test
				if err := keyDB.Insert(_testKey2, _testVal2); err != nil {
					log.Fatal(err.Error())
				}
				val, err := keyDB.Select(_testKey2)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(val).Equal(_testVal2)
			})

			g.It("should insert a key with a large value", func() {
				// function under test
				if err := keyDB.Insert(_testKey3, _testVal3); err != nil {
					log.Fatal(err.Error())
				}
				val, err := keyDB.Select(_testKey3)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(val).Equal(_testVal3)
			})

			g.It("should insert a key with an empty string value", func() {
				// function under test
				if err := keyDB.Insert(_testKey4, _testVal4); err != nil {
					log.Fatal(err.Error())
				}
				val4, err := keyDB.Select(_testKey4)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(val4).Equal(_testVal4)
			})
		})

		g.Describe("Select() Tests", func() {
			g.It("should assert error selecting non-existent key", func() {
				// function under test (selecting non-existent key)
				_, err := keyDB.Select("non_existent_key")
				g.Assert(err).Equal(errors.Errorf("Error getting value using key 'non_existent_key': redigo: nil returned"))
			})

			g.It("should assert error selecting non-existent key", func() {
				// function under test (selecting non-existent key)
				_, err := keyDB.Select("test_key 1")
				g.Assert(err).Equal(errors.Errorf("Error getting value using key 'test_key 1': redigo: nil returned"))
			})

			g.It("should assert error selecting non-existent key", func() {
				// function under test (selecting non-existent key)
				_, err := keyDB.Select(" test_key2")
				g.Assert(err).Equal(errors.Errorf("Error getting value using key ' test_key2': redigo: nil returned"))
			})

			g.It("should assert error selecting non-existent key", func() {
				// function under test (selecting non-existent key)
				_, err := keyDB.Select("test_key3 ")
				g.Assert(err).Equal(errors.Errorf("Error getting value using key 'test_key3 ': redigo: nil returned"))
			})
		})

		g.Describe("Delete() Tests", func() {
			g.It("should assert deleted key value pair does not exist", func() {
				// setup
				if err := keyDB.Insert(_testKey1, _testVal1); err != nil {
					log.Fatal(err.Error())
				}
				// setup assertion
				keyExists, err := keyDB.Exists(_testKey1)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
				// function under test
				err = keyDB.Delete(_testKey1)
				if err != nil {
					log.Fatalf(err.Error())
				}
				// final assertion
				keyExists, err = keyDB.Exists(_testKey1)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(false)
			})

			g.It("should assert deleted key value pair does not exist", func() {
				// setup
				if err := keyDB.Insert(_testKey2, _testVal2); err != nil {
					log.Fatal(err.Error())
				}
				// setup assertion
				keyExists, err := keyDB.Exists(_testKey2)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
				// function under test
				err = keyDB.Delete(_testKey2)
				if err != nil {
					log.Fatalf(err.Error())
				}
				keyExists, err = keyDB.Exists(_testKey2)
				if err != nil {
					log.Fatalf(err.Error())
				}
				// final assertion
				g.Assert(keyExists).Equal(false)
			})

			g.It("should assert deleted key value pair does not exist", func() {
				// setup
				if err := keyDB.Insert(_testKey3, _testVal3); err != nil {
					log.Fatal(err.Error())
				}
				// setup assertion
				keyExists, err := keyDB.Exists(_testKey3)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
				// function under test
				err = keyDB.Delete(_testKey3)
				if err != nil {
					log.Fatalf(err.Error())
				}
				// final assertion
				keyExists, err = keyDB.Exists(_testKey3)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(false)
			})

			g.It("should assert deleted key value pair does not exist", func() {
				// setup
				if err := keyDB.Insert(_testKey4, _testVal4); err != nil {
					log.Fatal(err.Error())
				}
				// setup assertion
				keyExists, err := keyDB.Exists(_testKey4)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
				// function under test
				err = keyDB.Delete(_testKey4)
				if err != nil {
					log.Fatalf(err.Error())
				}
				// final assertion
				keyExists, err = keyDB.Exists(_testKey4)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(false)
			})
		})

		g.Describe("Exists() Tests", func() {
			g.It("should assert non-existent key does not exist", func() {
				// function under test
				keyExists, err := keyDB.Exists("non_existent_key")
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(false)
			})

			g.It("inserted key value pair should exist", func() {
				// setup
				if err := keyDB.Insert(_testKey1, _testVal1); err != nil {
					log.Fatal(err.Error())
				}
				keyExists, err := keyDB.Exists(_testKey1)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
			})

			g.It("inserted key value pair should exist", func() {
				// setup
				if err := keyDB.Insert(_testKey2, _testVal2); err != nil {
					log.Fatal(err.Error())
				}
				keyExists, err := keyDB.Exists(_testKey2)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
			})

			g.It("inserted key value pair should exist", func() {
				// setup
				if err := keyDB.Insert(_testKey3, _testVal3); err != nil {
					log.Fatal(err.Error())
				}
				keyExists, err := keyDB.Exists(_testKey3)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
			})

			g.It("inserted key value pair should exist", func() {
				// setup
				if err := keyDB.Insert(_testKey4, _testVal4); err != nil {
					log.Fatal(err.Error())
				}
				keyExists, err := keyDB.Exists(_testKey4)
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyExists).Equal(true)
			})
		})

		g.Describe("CountKeys() Tests", func() {
			g.It("key count should be 0", func() {
				// function under test
				keyCount, err := keyDB.CountKeys()
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyCount).Equal(0)
			})

			g.It("key count should be 1", func() {
				// setup
				insertKeys(1, keyDB)
				// function under test
				keyCount, err := keyDB.CountKeys()
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyCount).Equal(1)
			})

			g.It("key count should be 2", func() {
				// setup
				insertKeys(2, keyDB)
				// function under test
				keyCount, err := keyDB.CountKeys()
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyCount).Equal(2)
			})

			g.It("key count should be 3", func() {
				// setup
				insertKeys(3, keyDB)
				// function under test
				keyCount, err := keyDB.CountKeys()
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyCount).Equal(3)
			})

			g.It("key count should be 4", func() {
				// setup
				insertKeys(4, keyDB)
				// function under test
				keyCount, err := keyDB.CountKeys()
				if err != nil {
					log.Fatalf(err.Error())
				}
				g.Assert(keyCount).Equal(4)
			})
		})

	})
}
