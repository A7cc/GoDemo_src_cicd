package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func searchHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// SQL注入漏洞
	query := "SELECT username FROM users WHERE id=" + id
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Fprintf(w, username)
	}
}

func initDB() {
	var err error
	db, err = sql.Open(
		"sqlite3",
		"./test.db",
	)
	if err != nil {
		panic(err)
	}
	db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER,
		username TEXT
	)
	`)
	db.Exec(`
	INSERT INTO users VALUES(1,'admin')
	`)
}

func main() {
	initDB()
	http.HandleFunc(
		"/search",
		searchHandler,
	)
	http.ListenAndServe(
		":8080",
		nil,
	)
}
