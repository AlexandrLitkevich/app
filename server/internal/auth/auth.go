package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = AuthUser{
	Username: "soso",
	Password: "123",
}

var mySigningKey = []byte("johenews")

// TODO сделать время жизни
type Response struct {
	StatusCode  int
	AccessToken string
	RefreshToken string
}

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST")


		var checkData AuthUser
		json.NewDecoder(r.Body).Decode(&checkData)
		// тут сравнение в бд
		if checkData.Username != user.Username || checkData.Password != user.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		handler.ServeHTTP(w, r)
	}
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Hour * 2160).Unix()

	tokenString, err := token.SignedString(mySigningKey) 
	
	if err != nil {
		logrus.Error(err)
	}

	return tokenString, nil
}

func Protected(w http.ResponseWriter, r *http.Request) {
	//Его можно изолировать если передать BasicAuth
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	validToken, err := GenerateJWT()
	fmt.Println(validToken)

	if err != nil {
		fmt.Println(err)
	}
	// dataBytes, _ := json.Marshal(Response{StatusCode: http.StatusOK, Status: "protected", Token: validToken})
	// (w).Write(dataBytes)
}
