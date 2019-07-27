package env

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func LoadEnv() {
	log.Info("Loading Env variable")
	LoadEnvByFile()

}

func LoadEnvByFile() {
	log.Info("Loading from .env file")

	err := godotenv.Load()

	if err != nil {
		log.Info(".env file not found. Basic env variable will be used")
	}
}

func FindEnvVarOrFail(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Env variable " + key + " not found")
	}
	return value
}

func FindEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Warning("Env variable " + key + " not found")
	}
	return value
}
