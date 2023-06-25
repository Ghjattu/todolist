package controller

import (
	"backend/utils"
	"testing"
)

func TestRegister(t *testing.T) {
	url := "http://127.0.0.1:8080/todolist/v1/user/register/?username=qly&password=123"
	method := "POST"
	utils.SendRequest(method, url, nil)
}
