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

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

type ResponseUser struct {
	Key      		int `json:"key"`
	Username 		string `json:"username" `
	Url      		string `json:"url"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func FromSQLite(conn *sql.DB) *SQLite {
	stmt, _ := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS users (
			key	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			url TEXT,
			password TEXT,
			accessToken TEXT,
			refreshToken TEXT
		);
	`)
	stmt.Exec()
	
	return &SQLite{
		DB: conn,
	}
}

// Функция для воводы данных для отображения
func (s *SQLite) Get() []ResponseUser {
	users := []ResponseUser{}
	rows, err := s.DB.Query("SELECT key,username,url FROM users")
	if err != nil {
		logrus.Error(err)
	}
	/* 
    	defer пишем после обработки ошибок во избежении паники
        Оператор задержки выполняеться после выполнения функции
     */
	defer rows.Close()
	var key int
	var username string
	var url string
	for rows.Next() {
		err := rows.Scan(&key, &username, &url)
		if err != nil {
			logrus.Error(err)
		}
		users = append(users, ResponseUser{
			Key: key,
			Username: username,
			Url: url})
	}
	// !!Всегда нужго проверять наличие ошибок поможет в отладке
	err = rows.Err()

	if err != nil {
		logrus.Error(err)
	}
	return users
}

func (s *SQLite) AuthUser(user AuthUser) bool {
	rows, err := s.DB.Query(`SELECT username,password FROM users WHERE username = ?`, user.Username)
	if err != nil {
		logrus.Error(err)
	}
	var (
		username string
		password string
	)
	for rows.Next() {

	if err := rows.Scan(&username, &password); err != nil {
		logrus.Error(err)
		}
	}
	// Проверяем есть ли такой пользователь
	if len(username) == 0 {
		return false
	}

	 return CheckPasswordHash(user.Password, password)
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

// TODO нужно разделить наверно
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
		// Мы используем встроенную функцию defer, чтобы отложить выполнение r.Body.Close ()
        defer r.Body.Close()
		if r.Method == "POST" {
			// Обрабатываем OPTION
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logrus.Error(err)
			}
			var user User
			json.Unmarshal(reqBody, &user)
			nameUnused := feed.Used(user.Username)
			// проверяем есть ли такое имя в базк данных
			if !nameUnused {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
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

func (s *SQLite) Used(name string) bool {
	rows, err := s.DB.Query("SELECT username FROM users WHERE username = ?", name)
	if err != nil {
		logrus.Error(err)
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		if err := rows.Scan(&username); err != nil {
			logrus.Error(err)
		}
   }
  logrus.Println("username",username == "")
   return username == ""
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
