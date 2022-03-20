package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/store"
	"strconv"
	"strings"
)

func DeleteUser(feed *store.SQLite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		if r.Method == "DELETE" {
			path := strings.Trim(r.URL.Path, "/")

			pathParts := strings.Split(path, "/")// возвращает массив сторок
			key, _ := strconv.Atoi(pathParts[1])// преобразуем в int

			//Запрос SQL
			stmt := `DELETE from users where key = ?`
			feed.DB.Exec(stmt, key)
			log.Println("User delete:" ,key)
			// Правельно ли переиспользовал???????
			newUsers := feed.Get()
			js, _ := json.Marshal(newUsers)
			w.Write(js)
		}
	}
}