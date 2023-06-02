package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db, derr = sql.Open("mysql", "user:pass@tcp(localhost:3306)/trpp")
)

func addUser(fn string, sn string, un string, pw string) {
	_, err := db.Query("INSERT INTO users (name, surname, username, password) VALUES (?, ?, ?, ?)", fn, sn, un, pw)
	if err != nil {
		panic(err)
	}
}

func addToFav(user string, favorites int) {
	result, _ := db.Query("SELECT user FROM favorites WHERE favorites = ? AND user = ?", favorites, user)
	var u string
	result.Scan(&u)

	if u == user {
		return
	}

	_, err := db.Query("INSERT INTO favorites (user, favorites) VALUES (?, ?)", user, favorites)
	if err != nil {
		panic(err)
	}
}

type fav struct {
	a int
}

func getAllFav(user string) []int {
	result, err := db.Query("SELECT favorites FROM favorites WHERE user = ?", user)
	if err != nil {
		panic(err)
	}

	l := []int{}

	for result.Next() {
		var a int
		result.Scan(&a)
		l = append(l, a)
	}
	// result.Scan(&l)

	return l
}

func main() {

	_, err := db.Query("CREATE TABLE IF NOT EXISTS users (id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, name varchar(45), surname varchar(45), username varchar(45), password varchar(45))")
	if err != nil {
		panic(err)
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS favorites (id int AI PK, user varchar(45) REFERENCES users(username), favorites int)")
	if err != nil {
		panic(err)
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))
	http.Handle("/sources/", http.StripPrefix("/sources/", http.FileServer(http.Dir("./public/sources"))))

	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/auth/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/auth/register", Register)

	http.HandleFunc("/", mainpage)
	http.HandleFunc("/addActor", addActor)
	http.HandleFunc("/contacts", contacts)
	http.HandleFunc("/discussion", discussion)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/kino", kino)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/profile", profile)

	http.HandleFunc("/film/", film)
	http.HandleFunc("/film/addtofav/", addFav)

	// parseTopAwait()

	// fmt.Println(parseActors("435"))

	// fmt.Println(getAllFav("user3"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
