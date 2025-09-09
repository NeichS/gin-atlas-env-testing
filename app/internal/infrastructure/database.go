package infrastructure

import (
	"fmt"
	"os"
)

type DbConn struct {
	Connstr string
}

func (d *DbConn) InitDatabase() *DbConn {
	host := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	usr := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPass := os.Getenv("POSTGRES_PASSWORD")

	connstr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbPort, usr, dbName, dbPass)

	return &DbConn{
		Connstr: connstr,
	}
}
