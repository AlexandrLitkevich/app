package store

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"github.com/sirupsen/logrus"
)


type SQLite struct {
	DB *sql.DB
}

type User struct {
	Key      		int `json:"key"`
	Username 		string `json:"username" `
	Url      		string `json:"url"`
	Password      	string `json:"password"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
}

type UserFront struct {
	Key      int 	`json:"key"`
	Username string `json:"username"`
	Url      string `json:"url"`
}
// Функция для воводы данных для отображения
func (s *SQLite) Get() []User {
	users := []User{}
	rows, err := s.DB.Query("SELECT * FROM users")
	if err != nil {
		logrus.Error(err)
	}
	defer rows.Close()
	var key int
	var username string
	var url string
	var password string
	var accessToken string
	var refreshToken string
	for rows.Next() {
		rows.Scan(&key, &username, &url, &password, &accessToken, &refreshToken)
		users = append(users, User{key, username, url, password, accessToken, refreshToken})
	}
	log.Println("users", users)
	return users
}

// FromSQLite creates a newfeed that uses sqlite
func FromSQLite(conn *sql.DB) *SQLite {
	stmt, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS
		users (
			key	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT
			url TEXT
			password TEXT
			accessToken TEXT
			refreshToken TEXT
		);
	`)
	stmt.Exec()
	return &SQLite{
		DB: conn,
	}
}


func UsersGet(users *SQLite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := users.Get()
		js, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Write(js)
	}
}

// TODO сделать обязательные поля
func CreateUser (feed *SQLite) http.HandlerFunc {
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

			var user User
			json.Unmarshal(reqBody, &user)
			user.Key = rand.Intn(1000000)
			user.Password = hashPassword(user.Password)	
			// Запрос SQL
			stmt := `INSERT INTO users(key, username, url, password)
			VALUES(?,?,?,?)`
			// Кладем в базу
			feed.DB.Exec(stmt, user.Key, user.Username, user.Url, user.Password)
			logrus.Info("create new user",user )

			newUsers := feed.Get()
			js, _ := json.Marshal(newUsers)
			w.Write(js)
		}
	}
}

func DeleteUser(feed *SQLite) http.HandlerFunc {
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
			
			newUsers := feed.Get()
			js, _ := json.Marshal(newUsers)
			w.Write(js)
		}
	}
}


func hashPassword(password string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		logrus.Error(err)
	}

    return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
