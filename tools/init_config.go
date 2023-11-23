package tools

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

var Config map[string]string

func GetEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("not found .env file, use default config")
	}
	Config = make(map[string]string)
	Config["DB_DRIVER"] = GetEnv("DB_DRIVER", "mysql")
	Config["DB_HOST"] = GetEnv("DB_HOST", "127.0.0.1")
	Config["DB_USER"] = GetEnv("DB_USER", "root")
	Config["DB_PASSWORD"] = GetEnv("DB_PASSWORD", "singularity")
	Config["DB_PORT"] = GetEnv("DB_PORT", "3306")
	Config["DB_NAME"] = GetEnv("DB_NAME", "test")
	Config["REDIS_HOST"] = GetEnv("REDIS_HOST", "127.0.0.1")
	Config["REDIS_PORT"] = GetEnv("REDIS_PORT", "6379")
	Config["REDIS_DB"] = GetEnv("REDIS_DB", "0")
	Config["REDIS_PASSWORD"] = GetEnv("REDIS_PASSWORD", "")
	Config["LOG_PATH"] = GetEnv("REDIS_PASSWORD", filepath.Join("data", "logs", "test.log"))
	fmt.Println(Config)
}
