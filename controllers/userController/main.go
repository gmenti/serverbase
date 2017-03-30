package userController

import (
	"net/http"
	"encoding/json"
)

type User struct {
	Name	string	`json:"name,omitempty"`
	Email	string	`json:"email,omitempty"`
}

func Index (writer http.ResponseWriter, _ *http.Request) {
	users := []User{
		{
			Name: "Giuseppe Menti",
			Email: "mentifg@gmail.com",
		},
	}

	json.NewEncoder(writer).Encode(users)
}