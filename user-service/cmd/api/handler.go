package main

import (
	"log"
	"net/http"

	"github.com/girigirig/user-service/cmd/jwt"
)

func (app *Config) LogIn(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.GenerateJwtToken()
	if err != nil {
		log.Println(err.Error())
	}

	if err = jwt.ValidateToken(token); err != nil {
		return
	}

	w.Write([]byte(token))
}

func (app *Config) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
