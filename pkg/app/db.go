package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/factotum/moneymaker/user-service/pkg/config"
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
			log.Println("Connecte to Postgres!")
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
