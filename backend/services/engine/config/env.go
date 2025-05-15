package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	LoggerLevel            string
	REDIS_URL              string
	IPWhiteLists           string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("HOST", ""),
		Port:                   getEnv("PORT", ""),
		DBUser:                 getEnv("DB_USER", ""),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName:                 getEnv("DB_NAME", ""),
		REDIS_URL:              getEnv("REDIS_URL", ""),
		LoggerLevel:            getEnv("LOG_LEVEL", "info"),
		IPWhiteLists:           getEnv("IP_WHITELISTS", ""),
		JWTSecret:              getEnv("JWT_SECRET", ""),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRY", 3600*24*7), // 7 days
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
