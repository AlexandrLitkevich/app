package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

)


var users []User

type userServer struct {
	store *userstore.UserStore
}

func NewUserServer() *userServer {
	store := userStore.New()
	return &userServer{store: store}
}


func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("create")
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
	log.Println("work delete")
	log.Println("method delete", r.Method)
	

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	// // Vars передаем туда запрос и получаем ключи
	// params := mux.Vars(r)
	// for index, item := range users {
	// 	if item.Key == params["id"] {
	// 		users = append(users[:index], users[index+1:]...)
	// 		break
	// 	}
	// }
	// json.NewEncoder(w).Encode(users)
}



func (us *userServer) userHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("work",req.Method)
}

func main() {
	// router := mux.NewRouter()
	// Мультиплексор поддерживат только точные пути
	mux := http.NewServeMux()
	server := NewUserServer()

	fs := http.FileServer(http.Dir("../client/build"))

	mux.Handle("/", fs)

	// users = append(users, User{Key: "1", Username: "Антон Иванов", Url: "www.ali.com"})
	// users = append(users, User{Key: "2", Username: "Гай Ричи", Url: "www.mail.com"})
	// users = append(users, User{Key: "6", Username: "Джон", Url: "www.google.com"})

	mux.HandleFunc("/users", getUsers)
	mux.HandleFunc("/added", createUser)
	// newServeMux не может обработать /user/{id}
	mux.HandleFunc("/user", server.userHandler)

	log.Println("Serving http://127.0.0.1:8000")



	// log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
	// log.Println("Serving http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", mux))

}
