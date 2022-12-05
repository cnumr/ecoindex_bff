package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ENV *Environment

type Environment struct {
	Env     string
	AppPort string
	ApiKey  string
	ApiUrl  string
}

func GetEnvironment() *Environment {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(currentPath + "/.env"); !os.IsNotExist(err) {
		fmt.Println("Load dotenv file")
		godotenv.Load()
	}

	return &Environment{
		Env:     getEnv("ENV", "dev"),
		AppPort: getEnv("APP_PORT", "3001"),
		ApiKey:  getEnv("API_KEY", ""),
		ApiUrl:  getEnv("API_URL", "https://ecoindex.p.rapidapi.com"),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
