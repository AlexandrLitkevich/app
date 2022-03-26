package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/store"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)



var user = store.AuthUser{
	Username: "soso",
	Password: "123",
}



// env или toml
var mySigningKey = []byte("johenews")

// TODO сделать время жизни
type Response struct {
	StatusCode  int
	AccessToken string
}


func BasicAuth(feed *store.SQLite,handler http.HandlerFunc) http.HandlerFunc {
	type Fail struct {
		Status int `json:"status"`
		Desc string `json:"desc"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST")


		var checkData store.AuthUser
		json.NewDecoder(r.Body).Decode(&checkData)
		isSingIn := feed.AuthUser(checkData)
		// тут сравнение в бд
		if !isSingIn {
			w.WriteHeader(http.StatusUnauthorized)
			responseBytes, _ := json.Marshal( Fail{ Status: http.StatusUnauthorized, Desc: "Данные не верны"} ) 
			w.Write(responseBytes)
			return
		}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	validToken, err := GenerateJWT()
	fmt.Println(validToken)

	if err != nil {
		fmt.Println(err)
	}

	dataBytes, _ := json.Marshal(Response{StatusCode: http.StatusOK, AccessToken: validToken})
	w.Write(dataBytes)
}
