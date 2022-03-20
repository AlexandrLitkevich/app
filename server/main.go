package main

import (
	"database/sql"
	"log"
	"net/http"
	"server/handler"
	"server/internal/store"
	"server/internal/userstore"

	_ "github.com/mattn/go-sqlite3"
)


type userServer struct {
	store *userstore.UserStore
}

func NewUserServer() *userServer {
	store := userstore.New()
	return &userServer{store: store}
}

/* 
1) Как парвельно работать с данными из ответа?(Работа с JSON)
2) Основы  SQL и работы с BD
3) Основные паттерны и алгоритмы на беке
4) как не беке обрабатывают OPTIONS???
5) Переменная в виде одной буквы это норм?
6) Требуемые базовые пакеты?
7) Создание бд
8) Вызов функций в соедних файлах(users_get & delete_users)
*/

func main() {
	database, _ := sql.Open("sqlite3", "./users.db")

	err := database.Ping()
	if err != nil {
		log.Panicln(err)
	}

	feed := store.FromSQLite(database)

	// Мультиплексор поддерживат только точные пути
	mux := http.NewServeMux()
	//Ping ?
	fs := http.FileServer(http.Dir("../client/build"))

	mux.Handle("/", fs)

	mux.HandleFunc("/users", handler.UsersGet(feed))
	mux.HandleFunc("/added", handler.AddedUser(feed))
	mux.HandleFunc("/user/", handler.DeleteUser(feed))

	log.Println("Serving http://127.0.0.1:8000")

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", mux))

}
