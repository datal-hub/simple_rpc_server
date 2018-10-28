package dbtest

import (
	"database/sql"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"rpc-server/models"
)

type DbTest struct {
	DB       *sql.DB
	Mock     sqlmock.Sqlmock
	TestUser models.User
}

func (db *DbTest) SqlDB() *sql.DB {
	return db.DB
}

func (db *DbTest) Close() {
	db.DB.Close()
}
