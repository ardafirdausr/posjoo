package driver

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQLDB(DBURI string) (*sql.DB, error) {
	DB, err := sql.Open("mysql", DBURI+"?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return DB, nil
}
