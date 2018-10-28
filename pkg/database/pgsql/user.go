package pgsql

import (
	"errors"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"rpc-server/models"
)

func (db *PgSQL) Add(user *models.User) error {
	rdb := reform.NewDB(db.DB, postgresql.Dialect, nil)
	if err := rdb.FindByPrimaryKeyTo(&models.User{}, user.Uuid); err == nil {
		return errors.New("user already exists")
	}
	err := rdb.Save(user)
	return err
}

func (db *PgSQL) Get(uuid string) (models.User, error) {
	rdb := reform.NewDB(db.DB, postgresql.Dialect, nil)
	user := models.User{}
	if err := rdb.FindByPrimaryKeyTo(&user, uuid); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (db *PgSQL) Delete(uuid string) error {
	rdb := reform.NewDB(db.DB, postgresql.Dialect, nil)
	_, err := rdb.DeleteFrom(models.UserTable, " WHERE uuid = $1", uuid)
	return err
}
