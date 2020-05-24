package config

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestWithProdEnvArg(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(validEnvArg("prod"), true, "should be a valid env arg")
}

func TestWithDevEnvArg(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(validEnvArg("dev"), true, "should be a valid env arg")
}

func TestWithTestEnvArg(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(validEnvArg("test"), true, "should be a valid env arg")
}

func TestWithSpaceInProdEnvArg(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(validEnvArg("prod "), false, "should be a valid env arg")
}

func TestWithSpaceInTestEnvArg(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(validEnvArg(" test"), false, "should be a valid env arg")
}

func TestWithSpaceInDevEnvArg(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(validEnvArg(" dev"), false, "should be a valid env arg")
}

func TestWithInvalidEnvArg(t *testing.T) {
  assert := assert.New(t)
  assert.Equal(validEnvArg("invalid"), false, "should be a valid env arg")
}
