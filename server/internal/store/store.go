package store

import (
	"database/sql"
)


type SQLite struct {
	DB *sql.DB
}

type User struct {
	Key      int `json:"key"`
	Username string `json:"username"`
	Url      string `json:"url"`
}




func (s *SQLite) Get() []User {
	users := []User{}
	rows, _ := s.DB.Query("SELECT * FROM users")
	defer rows.Close()
	var key int
	var username string
	var url string
	for rows.Next() {
		rows.Scan(&key, &username, &url)
		users = append(users, User{key, username, url})
	}
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
		);
	`)
	stmt.Exec()
	return &SQLite{
		DB: conn,
	}
}
