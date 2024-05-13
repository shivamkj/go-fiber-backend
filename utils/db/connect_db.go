package db

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

func connectDb(driverName string, config DbConfig) *sql.DB {
	db, err := sql.Open(driverName, config.ConnectionURL)
	if err != nil {
		panic("database connection failed, err:" + err.Error())
	}

	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(int(time.Minute) * config.ConnectionMaxLifetime))
	db.SetConnMaxIdleTime(time.Duration(int(time.Minute) * config.ConnectionMaxIdleTime))

	if err := db.Ping(); err != nil {
		defer db.Close()
		panic("can not ping to database, err:" + err.Error())
	}

	return db
}
