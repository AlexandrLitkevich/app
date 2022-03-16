package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// использовать id а не key
type User struct {
	Key      string `json:"key"`
	Username string `json:"username"`
	Url      string `json:"url"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// До этого идет OPTION поэтому 2 раза сробатывает ф-ия
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	reqBody, _ := ioutil.ReadAll(r.Body)
	// проверка если придет пустой body
	if len(reqBody) != 0 {
		var user User

		json.Unmarshal(reqBody, &user)
		user.Key = strconv.Itoa(rand.Intn(1000000))
		users = append(users, user)
		json.NewEncoder(w).Encode(users)
	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// Vars передаем туда запрос и получаем ключи
	params := mux.Vars(r)
	for index, item := range users {
		if item.Key == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()
	users = append(users, User{Key: "1", Username: "Антон Иванов", Url: "www.ali.com"})
	users = append(users, User{Key: "2", Username: "Гай Ричи", Url: "www.mail.com"})
	users = append(users, User{Key: "6", Username: "Джон", Url: "www.google.com"})

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/added", createUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8000", r))
}
