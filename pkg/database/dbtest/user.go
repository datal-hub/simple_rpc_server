package dbtest

import (
	"gopkg.in/reform.v1"

	"rpc-server/models"
)

func (db *DbTest) Add(user *models.User) error {
	return nil
}

func (db *DbTest) Get(uuid string) (models.User, error) {
	if uuid == db.TestUser.Uuid {
		return db.TestUser, nil
	}
	return models.User{}, reform.ErrNoRows
}

func (db *DbTest) Delete(uuid string) error {
	return nil
}
