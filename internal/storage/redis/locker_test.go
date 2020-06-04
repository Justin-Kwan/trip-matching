package redis

import (
  "log"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"order-matching/internal/config"
)

var (
	_locker *Locker
)

// type lockerTestConstants struct {
//
// }

func setupLockerTests() func() {
	tc := newGeoDBTestConstants()

	cfg, _ := config.NewConfig(tc.configFilePath, tc.env)
	geoDBPool := NewPool(&(*cfg).RedisGeoDB)

	_geoDB = NewGeoDB(geoDBPool, tc.setIndex)
	_geoDB.Clear()

	_locker = NewLocker(geoDBPool)

	return func() {
		_geoDB.Clear()
	}
}

func TestLockKey(t *testing.T) {
	teardownTests := setupLockerTests()
	defer teardownTests()

	var testCases = []struct {
		inputKeyId  string
		inputTTL    int
		expectedErr error
	}{
		{"test_key_id1", -1, errors.Errorf("ERR invalid expire time in set")},
		{"test_key_id1", 50, nil},
		{"test_key_id1", 50, errors.Errorf("Error, key is already locked")},
		{"test_key_id2", 50, nil},
		{"test_key_id2", 50, errors.Errorf("Error, key is already locked")},
		{"test_key_id2", 25, errors.Errorf("Error, key is already locked")},
		{"test_key_id2", 0, errors.Errorf("ERR invalid expire time in set")},
	}

	for _, testCase := range testCases {
    // function under test
		err := _locker.LockKey(testCase.inputKeyId, testCase.inputTTL)

		if testCase.expectedErr != nil {
			assert.EqualError(t, err, testCase.expectedErr.Error(), "should assert error is returned")
		} else {
			assert.NoError(t, err, "should assert no error is returned")
		}
	}
}

func TestUnlockKey(t *testing.T) {
  teardownTests := setupLockerTests()
  defer teardownTests()

  err := _locker.UnlockKey("non_existent_keyid")
  assert.EqualError(t, err, "Error unlocking key 'non_existent_keyid'", "should assert error is returned")

  var testCases = []struct {
		inputKeyId  string
		inputTTL    int
		expectedErr error
	}{
		{"test_key_id1", 50, nil},
		{"test_key_id2", 50, nil},
	}

	for _, testCase := range testCases {
    // setup
		err := _locker.LockKey(testCase.inputKeyId, testCase.inputTTL)
    if err != nil {
      teardownTests()
      log.Fatalf(err.Error())
    }
    
    // function under test
    err = _locker.UnlockKey(testCase.inputKeyId)
    assert.NoError(t, err, "should assert no error is returned")
	}
}
