package database

import (
	"database/sql"
	"fmt"

	"rpc-server/models"
	"rpc-server/pkg/database/pgsql"
	"rpc-server/pkg/settings"
)

// DB define interface for work with database
type DB interface {
	IsEmpty() bool
	Clear()
	Init(force bool) error
	Close()
	SqlDB() *sql.DB

	Add(user *models.User) error
	Get(uuid string) (models.User, error)
	Delete(uuid string) error
}

func NewDB() (DB, error) {
	if Testing == true {
		return testDb()
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=disable",
		settings.DB.Host, settings.DB.Port, settings.DB.Password, settings.DB.User, settings.DB.Database))
	return &pgsql.PgSQL{DB: db}, err
}
