package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ENV *Environment

type Environment struct {
	Env         string
	AppPort     string
	AppUrl      string
	ApiKey      string
	ApiUrl      string
	CacheTtl    string
	EcoindexUrl string
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
		Env:         getEnv("ENV", "dev"),
		AppPort:     getEnv("APP_PORT", "3001"),
		AppUrl:      getEnv("APP_URL", "http://localhost:3001"),
		ApiKey:      getEnv("API_KEY", ""),
		ApiUrl:      getEnv("API_URL", "https://ecoindex.p.rapidapi.com"),
		CacheTtl:    getEnv("CACHE_TTL", fmt.Sprintf("%d", 60*60*24*7)),
		EcoindexUrl: getEnv("ECOINDEX_URL", "https://www.ecoindex.fr"),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
