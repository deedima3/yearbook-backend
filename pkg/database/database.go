package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB) {
	dataSourceName := "zh9b94e86tl6:pscale_pw_ftOx0YKNFJKmFaWI-gjNtyReEfUubj-QwYfrLcI_PiE@tcp(ddzkodr24h6n.ap-southeast-2.psdb.cloud)/yearbook_db?tls=true"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("ERROR GetDatabase sql open connection fatal error: %v\n", err)
		for{
			log.Printf("INFO GetDatabase re-attempting to reconnect to database...")
			time.Sleep(1 * time.Second)
			if err == nil {
				break
			}
		}
	}
	if db.Ping(); err != nil {
		log.Printf("ERROR GetDatabase ping fatal error: %v\n", err)
		for{
			log.Println("INFO GetDatabase re-attempting to reconnect to database...")
			time.Sleep(1 * time.Second)
			db, err := sql.Open("mysql", dataSourceName)
			err2 := db.Ping()
			if err == nil && err2 == nil {
				break
			}
		}
	}
	log.Printf("INFO GetDatabase database connection: established successfully with %s\n", dataSourceName)
	return db
}

func TestConnection() {
	db:= GetDatabase()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
}
