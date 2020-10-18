package config

import (
	"os"
)

func GetEnv() string {
	if env := os.Getenv("ENV"); env != "" {
		return env
	}
	return "dev"
}
