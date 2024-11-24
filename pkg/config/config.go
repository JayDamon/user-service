package config

import (
	"fmt"
	"github.com/jaydamon/moneymakergocloak"
	"github.com/jaydamon/moneymakerrabbit"
	"log"
	"os"
	"strconv"
)

type Config struct {
	HostPort       string
	ConfigureCors  bool
	DB             *DBConfig
	KeyCloakConfig *moneymakergocloak.Configuration
	Rabbit         *moneymakerrabbit.Configuration
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

	configureCors := getOrDefaultBool("CONFIGURE_CORS", true)

	return &Config{
		HostPort:       "8091",
		ConfigureCors:  configureCors,
		DB:             configureDB(),
		KeyCloakConfig: moneymakergocloak.NewConfiguration(),
		Rabbit:         moneymakerrabbit.NewConfiguration(),
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
	val := os.Getenv(envVar)
	if val == "" {
		return defaultVal
	}
	return val
}

func getOrDefaultBool(envVar string, defaultVal bool) bool {
	val := os.Getenv(envVar)
	var returnVal = defaultVal
	if val == "true" {
		returnVal = true
	} else if val == "false" {
		returnVal = false
	}

	return returnVal
}
