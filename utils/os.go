package utils

import (
	"os"
	"strconv"
)

func GetEnvDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func GetIntEnvDefault(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	res, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return res
}
