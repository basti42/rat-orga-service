package utils

import (
	"log"
	"os"
)

func mustEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("missing env variable '%v'", key)
	}
	return value
}

var (
	SERVICE_NAME = mustEnv("SERVICE_NAME")
	PORT         = mustEnv("PORT")
	DB_HOST      = mustEnv("DB_HOST")
	DB_PORT      = mustEnv("DB_PORT")
	DB_NAME      = mustEnv("DB_NAME")
	DB_USER      = mustEnv("DB_USER")
	DB_PASSWORD  = mustEnv("DB_PASSWORD")
	JWT_SECRET   = mustEnv("JWT_SECRET")
)
