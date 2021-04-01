package util

import (
	"log"
	"os"
)

func GetEnv(key, fallback string) string {
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
