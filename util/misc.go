package util

import (
	"log"
	"os"
	"strconv"
	"time"
)

type TimeSlice []time.Time

func (s TimeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s TimeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TimeSlice) Len() int           { return len(s) }

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
