package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/factotum/moneymaker/user-service/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var counts int64

func connectToDB(config *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.Name)

	log.Printf("Connecting to DB: %s\n", psqlInfo)

	for {
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return db
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}

func performDbMigration(db *sql.DB, config *config.Config) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		config.DB.Name, driver)
	if err != nil {
		log.Panic(err)
	}
	m.Up()
}
