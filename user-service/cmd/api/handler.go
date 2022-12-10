package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("SECRET"))

func (app *Config) LogIn(w http.ResponseWriter, r *http.Request) {
	token, err := generateJwtToken()
	if err != nil {
		log.Println(err.Error())
	}

	w.Write([]byte(token))
}

func generateJwtToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "Gideon"

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func verifyJWT(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	
}