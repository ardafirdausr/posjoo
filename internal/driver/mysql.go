package driver

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQLDB(DBURI string) (*sql.DB, error) {
	// u, err := url.Parse(DBURI)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// q := u.
	// q.Set("parseTime", "true")
	// u.RawQuery = q.Encode()

	// DBURI = u.String()
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
