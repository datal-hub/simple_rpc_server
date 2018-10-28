package testing

import (
	"time"

	"rpc-server/models"
)

var TestUser = models.User{
	Uuid:  "d85e28e8-997a-4805-8703-e3e7913ad608",
	Login: "test",
	Dttm:  time.Date(2018, 10, 20, 1, 2, 3, 0, time.UTC),
}

var TestNotExistUser = models.User{
	Uuid:  "e1e7cfdc-da83-11e8-9f8b-f2801f1b9fd1",
	Login: "new",
	Dttm:  time.Date(2018, 10, 20, 1, 2, 3, 0, time.UTC),
}
