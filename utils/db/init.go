package db

import (
	"database/sql"
	"strconv"
)

const (
	Postgres = iota + 1
	MySQL
	SQLite
)

type DbConfig struct {
	Type                  int    `yaml:"type"`
	ConnectionURL         string `yaml:"connection_url"`
	MaxOpenConnections    int    `yaml:"maxOpenConnections"`
	MaxIdleConnections    int    `yaml:"maxIdleConnections"`
	ConnectionMaxLifetime int    `yaml:"connectionMaxLifetime"`
	ConnectionMaxIdleTime int    `yaml:"connectionMaxIdleTime"`
}

var dbInit = false

func InitDB(config DbConfig) *sql.DB {
	if dbInit {
		panic("database already initialised")
	}

	if strconv.IntSize != 64 {
		panic("unsupported OS architecture")
	}

	var db *sql.DB
	switch config.Type {
	case Postgres:
		db = connectDb("pgx", config)
	case MySQL:
		db = connectDb("mysql", config)
	case SQLite:
		config.ConnectionURL += "?_journal=WAL&_timeout=5000&_fk=true"
		db = connectDb("sqlite3", config)
	default:
		panic("unsupported database type passed in config")
	}

	dbInit = true
	return db
}
