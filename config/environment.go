package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var ENV *Environment

type Environment struct {
	Env          string
	AppPort      string
	AppUrl       string
	ApiKey       string
	ApiUrl       string
	BadgeVersion string
	CacheDsn     string
	CacheEnabled bool
	CacheTtl     int
	CDNUrl       string
	EcoindexUrl  string
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

	ttlInt, err := strconv.Atoi(getEnv("CACHE_TTL", fmt.Sprintf("%d", 60*60*24*7)))
	if err != nil {
		panic(err)
	}

	cacheEnabled, err := strconv.ParseBool(getEnv("CACHE_ENABLED", "true"))
	if err != nil {
		panic(err)
	}

	return &Environment{
		Env:          getEnv("ENV", "dev"),
		AppPort:      getEnv("APP_PORT", "3001"),
		AppUrl:       getEnv("APP_URL", "http://localhost:3001"),
		ApiKey:       getEnv("API_KEY", ""),
		ApiUrl:       getEnv("API_URL", "https://ecoindex.p.rapidapi.com"),
		BadgeVersion: getEnv("BADGE_VERSION", "v1"),
		CacheDsn:     getEnv("CACHE_DSN", "localhost:6379"),
		CacheEnabled: cacheEnabled,
		CacheTtl:     ttlInt,
		CDNUrl:       getEnv("CDN_URL", "https://cdn.jsdelivr.net/gh/cnumr/ecoindex_badge"),
		EcoindexUrl:  getEnv("ECOINDEX_URL", "https://www.ecoindex.fr"),
	}
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
