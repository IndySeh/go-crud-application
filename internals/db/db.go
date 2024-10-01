package db

import (
	"errors"
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// missing database connection error
	if dbDriver == "" || dbUser == "" || dbPass == "" || dbName == "" {
		return nil, errors.New("database connection parameters are missing")
	}

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, errors.New(err.Error())
	}

	return db, nil
}
