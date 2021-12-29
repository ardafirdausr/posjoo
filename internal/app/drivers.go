package app

import (
	"database/sql"
	"os"

	"github.com/ardafirdausr/posjoo-server/internal/driver"
)

type drivers struct {
	MySQL *sql.DB
}

func newDrivers() (*drivers, error) {
	drivers := new(drivers)

	MySQLURI := os.Getenv("MYSQL_URI")
	MySQLDB, err := driver.ConnectToMySQLDB(MySQLURI)
	if err != nil {
		return nil, err
	}
	drivers.MySQL = MySQLDB

	return drivers, nil
}
