package db

import (
	"database/sql"
	"fmt"

	"github.com/qnify/api-server/utils/fiber"
)

type DbX struct {
	*sql.DB
	dbType int
}

func NewDBX(dbType int, db *sql.DB) *DbX {
	return &DbX{
		dbType: dbType,
		DB:     db,
	}
}

func (db *DbX) QueryRowX(query string, args ...any) *sql.Row {
	switch db.dbType {
	case Postgres, SQLite:
		return db.QueryRow(query, args...)
	}
	panic("invalid db type passed in QueryRowX")
}

func (db *DbX) ExecX(pgQuery string, msQuery string, args ...any) (sql.Result, error) {
	switch db.dbType {
	case Postgres, SQLite:
		return db.Exec(pgQuery, args...)
	case MySQL:
		return db.Exec(msQuery, args...)
	}
	panic("invalid db type passed in ExecX")
}

func (db *DbX) Listx(query string, filter *fiber.QueryFilters) (*sql.Rows, error) {
	switch db.dbType {
	case Postgres, SQLite:
		qb := NewQuery(filter, PgParam)
		return db.Query(query+qb.Query(), qb.Params()...)
	case MySQL:
		qb := NewQuery(filter, MsParam)
		return db.Query(query+qb.Query(), qb.Params()...)
	}
	panic("invalid db type passed in Listx")
}

func (db *DbX) InsertX(pgQuery string, msQuery string, args ...any) (int, error) {
	switch db.dbType {
	case Postgres, SQLite:
		var lastInsertID int
		err := db.QueryRow(pgQuery, args...).Scan(&lastInsertID)
		return lastInsertID, err
	case MySQL:
		fmt.Println(args...)
		result, err := db.Exec(msQuery, args...)
		if err != nil {
			return 0, err
		}
		lastInsertID, err := result.LastInsertId()
		return int(lastInsertID), err
	}
	panic("invalid db type passed in InsertX")
}
