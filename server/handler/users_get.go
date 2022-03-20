package handler

import (
	"encoding/json"
	"net/http"
	"server/internal/store"
)

func UsersGet(users *store.SQLite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := users.Get()
		js, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Write(js)
	}
}
