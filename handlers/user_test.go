package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/reform.v1"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/stretchr/testify/assert"

	. "rpc-server/handlers"
	"rpc-server/models"
	testData "rpc-server/models/testing"
	"rpc-server/pkg/database"
)

func execute(t *testing.T, s *rpc.Server, method string, req, res interface{}) error {
	if !s.HasMethod(method) {
		t.Fatal("Expected to be registered:", method)
	}

	buf, _ := json.EncodeClientRequest(method, req)
	body := bytes.NewBuffer(buf)
	r, _ := http.NewRequest("POST", "http://localhost:8080/", body)
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)

	return json.DecodeClientResponse(w.Body, res)
}

func TestUserApiGetExist(t *testing.T) {
	database.Testing = true

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(UserApi), "")

	var res *models.User
	args := testData.TestUser.Uuid
	if err := execute(t, s, "UserApi.Get", &args, &res); err != nil {
		t.Error("Expected err to be nil, but got:", err)
	}
	assert.Equal(t, testData.TestUser.Login, res.Login)
}

func TestUserApiGetNotExist(t *testing.T) {
	database.Testing = true

	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(UserApi), "")

	var res *models.User
	args := testData.TestNotExistUser.Uuid
	err := execute(t, s, "UserApi.Get", &args, &res)
	assert.Equal(t, err, reform.ErrNoRows)
}
