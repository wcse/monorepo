package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	env := GetEnv()
	assert.Equal(t, env, "dev")

	_ = os.Setenv("ENV", "stg")
	env = GetEnv()
	assert.Equal(t, env, "stg")

	_ = os.Setenv("ENV", "beta")
	env = GetEnv()
	assert.Equal(t, env, "beta")
}
