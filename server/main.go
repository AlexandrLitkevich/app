package main

import (
	"database/sql"
	"log"
	"net/http"
	"server/internal/auth"
	"server/internal/store"

	_ "github.com/mattn/go-sqlite3"
)

/*
1) Как парвельно работать с данными из ответа?(Работа с JSON) десорилизация (формат в байтах)
2) Основы  SQL и работы с BD
3) Основные паттерны и алгоритмы на беке
4) как не беке обрабатывают OPTIONS???
5) для чего пишут api ???
*/

func main() {
	database, _ := sql.Open("sqlite3", "./users.db")
	// Проверка подключения
	err := database.Ping()

	if err != nil {
		log.Panicln(err)
	}

	feed := store.FromSQLite(database)

	// Мультиплексор поддерживат только точные пути
	mux := http.NewServeMux()
	
	fs := http.FileServer(http.Dir("../client/build"))

	mux.Handle("/", fs)

	mux.HandleFunc("/users", store.UsersGet(feed))
	mux.HandleFunc("/added", store.CreateUser(feed))
	mux.HandleFunc("/user/", store.DeleteUser(feed))
	mux.HandleFunc("/api/auth/", auth.BasicAuth(auth.Protected))

	log.Println("Serving http://127.0.0.1:8000")

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", mux))

}
