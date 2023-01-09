package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	HostPort string
	DB       *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		HostPort: "8091",
		DB:       configureDB(),
	}
}

func configureDB() *DBConfig {

	host := getOrDefault("DB_HOST", "localhost")
	strPort := getOrDefault("DB_PORT", "5432")
	port, err := strconv.Atoi(strPort)
	if err != nil {
		log.Panic(fmt.Printf("Port %s type is incorrect, must be int", strPort))
	}
	user := getOrDefault("DB_USER", "postgres")
	password := getOrDefault("DB_PASSWORD", "password")
	dbname := getOrDefault("DB_NAME", "users")
	charset := getOrDefault("DB_CHARSET", "utf8")

	return &DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     dbname,
		Charset:  charset,
	}
}

func getOrDefault(envVar string, defaultVal string) string {
	val := os.Getenv("DB_PORT")
	if val == "" {
		return defaultVal
	}
	return val
}
