package jwt

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Jwtx struct{}

var secretKey = []byte(os.Getenv("SECRET"))

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJwtToken() (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &JWTClaim{
		Email:    "fakeuser@gmail.com",
		Username: "fakeuser@gmail.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("issued occured while parsing claim")
	}

	claim, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't find claim")
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return
}

func Authorization(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := ValidateToken(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	}
}
