package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "2zlj4cmjec4v:pscale_pw_pSddB2ehFi8I0Bf3plUoeG2NWp2t6HMuPmuyg8BgDAw@tcp(ddzkodr24h6n.ap-southeast-2.psdb.cloud)/yearbook_db?tls=true")
	return db, err
}

func TestConnection() {
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
}
