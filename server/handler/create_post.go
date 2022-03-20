package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/internal/store"
	"database/sql"
	"math/rand"
)

/* 
	Что происходит когда мы пишем
	 func (conn *store.SQLite) CreateUser() http.HandlerFunc
	 (conn *store.SQLite) это же не аргументы
*/

type SQLite struct {
	DB *sql.DB
}


func AddedUser (feed *store.SQLite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		if r.Body == nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        defer r.Body.Close()

		if r.Method == "POST" {
			// Обрабатываем OPTION
			reqBody, _ := ioutil.ReadAll(r.Body)

			var user store.User
			json.Unmarshal(reqBody, &user)
			user.Key = rand.Intn(1000000)

			// Запрос SQL
			stmt := `INSERT INTO users (key, username, url)
			VALUES(?,?,?)`
			// Кладем в базу
			feed.DB.Exec(stmt, user.Key, user.Username, user.Url)
			log.Println("new user",user )

			// Правельно ли переиспользовал???????
			newUsers := feed.Get()
			js, _ := json.Marshal(newUsers)
			w.Write(js)
		}
	}
}
