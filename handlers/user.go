package handlers

import (
	"net/http"

	"rpc-server/models"
	"rpc-server/pkg/database"
)

type UserApi struct{}

func (UserApi) Add(r *http.Request, args *string, resp *string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	login := *args
	newUser, err := models.NewUser(login)
	if err != nil {
		return err
	}
	if err := db.Add(&newUser); err != nil {
		return err
	}
	*resp = newUser.Uuid
	return nil
}

func (UserApi) Get(r *http.Request, args *string, resp *models.User) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	*resp, err = db.Get(*args)
	if err != nil {
		return err
	}
	return nil
}

func (UserApi) Delete(r *http.Request, args *string, resp *string) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	if err := db.Delete(*args); err != nil {
		return err
	}
	resp = args
	return nil
}
