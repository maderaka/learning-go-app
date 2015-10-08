package core

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DbName     = "go_member"
	DbUsername = "rakateja"
	DbPassword = "password"
	DbHost     = "localhost"
	SslMode    = "disable"
	DriverName = "postgres"
)

func NewDB(dataSource string) (*sql.DB, error) {
	db, err := sql.Open(DriverName, dataSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
