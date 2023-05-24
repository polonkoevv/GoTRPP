package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db, derr = sql.Open("mysql", "root:1234@tcp(localhost:3306)/trpp")
)

func addUser(fn string, sn string, un string, pw string) {
	_, err := db.Query("INSERT INTO users (name, surname, username, password) VALUES (?, ?, ?, ?)", fn, sn, un, pw)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	http.Handle("/sources/", http.StripPrefix("/sources/", http.FileServer(http.Dir("./public/sources"))))

	http.HandleFunc("/auth/signin", Signin)
	http.HandleFunc("/auth/welcome", Welcome)
	http.HandleFunc("/auth/refresh", Refresh)
	http.HandleFunc("/auth/logout", Logout)
	http.HandleFunc("/auth/register", Register)

	http.HandleFunc("/", mainpage)
	http.HandleFunc("/addActor", addActor)
	http.HandleFunc("/contacts", contacts)
	http.HandleFunc("/discussion", discussion)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/kino", kino)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
