package storage

import (
	"database/sql"
	"log"

	config "github.com/kaushikkumarbora/TurnedIn/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDB(params ...string) *sql.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = sql.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *sql.DB {
	return DB
}
