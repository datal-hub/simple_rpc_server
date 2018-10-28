package models

import (
	"time"

	"github.com/satori/go.uuid"
)

//go:generate reform

// reform:users
type User struct {
	Uuid  string    `reform:"uuid,pk" json:"uuid"`
	Login string    `reform:"login" json:"login"`
	Dttm  time.Time `reform:"dttm" json:"dttm"`
}

func NewUser(login string) (User, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return User{}, nil
	}
	return User{
		Uuid:  uid.String(),
		Login: login,
		Dttm:  time.Now().UTC(),
	}, nil
}
