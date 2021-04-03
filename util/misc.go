package util

import (
	"log"
	"os"
	"strconv"
)

func GetEnvString(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		if len(fallback) > 0 {
			return fallback
		} else {
			log.Fatalf("mandatory environment variable missing: %s", key)
		}
	}
	return value
}

func GetEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	r, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("environment variable %s is not set to an integer", key)
	}
	return r
}
