package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Env = initEnv()

type Config struct {
	Port        string
	DatabaseUrl string
}

func initEnv() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseUrl: getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, d string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return d
}
