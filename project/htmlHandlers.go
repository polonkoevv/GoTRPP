package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func mainpage(w http.ResponseWriter, r *http.Request) {

	var path = "./public/index.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}

func addActor(w http.ResponseWriter, r *http.Request) {

	var path = "./public/addActor.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}
func contacts(w http.ResponseWriter, r *http.Request) {

	var path = "./public/contacts.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}
func discussion(w http.ResponseWriter, r *http.Request) {

	var path = "./public/discussion.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}
func gallery(w http.ResponseWriter, r *http.Request) {

	var path = "./public/gallery.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}

func kino(w http.ResponseWriter, r *http.Request) {

	// ts, err := template.ParseFiles("./public/kino.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }

	var path = "./public/kino.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, map[string]interface{}{"films": parseTopAwait()["films"]})
}

func film(w http.ResponseWriter, r *http.Request) {

	// ts, err := template.ParseFiles("./public/kino.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }

	// w.Header().Get()

	r.ParseForm()
	// они все тут
	params := r.Form

	fmt.Println(params["id"])

	var path = "./public/film.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, map[string]interface{}{"actors": parseActors(params["id"][0])[0:10],
		"posterUrl":   parseFilm(params["id"][0])["posterUrl"],
		"description": parseFilm(params["id"][0])["description"],
		"id":          params["id"][0],
		"genres":      parseFilm(params["id"][0])["genres"]})
}

func login(w http.ResponseWriter, r *http.Request) {

	var path = "./public/login.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}
func register(w http.ResponseWriter, r *http.Request) {

	var path = "./public/register.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}

func addFav(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := r.Form

	fmt.Println(params["id"])

	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			fmt.Println(1)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		fmt.Println(2)

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// We then get the name of the user from our session map, where we set the session token
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		fmt.Println(3)
		w.WriteHeader(http.StatusUnauthorized)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	i, _ := strconv.Atoi(params["id"][0])

	addToFav(userSession.username, i)
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func profile(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			fmt.Println(1)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		fmt.Println(2)

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// We then get the name of the user from our session map, where we set the session token
	userSession, exists := sessions[sessionToken]
	if !exists {
		// If the session token is not present in session map, return an unauthorized error
		fmt.Println(3)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	flms := []interface{}{}

	mp := map[interface{}]int{}

	ids := getAllFav(userSession.username)
	// var i int = 0
	for i := 0; i < len(ids); i++ {
		num := strconv.FormatInt(int64(ids[i]), 10)
		flms = append(flms, parseFilm(num)["nameRu"])
		mp[parseFilm(num)["nameRu"]] = ids[i]
	}

	fmt.Println(flms)

	var path = "./public/profile.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, map[string]interface{}{
		"films": mp,
	})
}

func actors(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// pars := r.Form

	// templateParser, err := template.ParseFiles("./pubilc/templates/account.html")
	// if err != nil{
	// 	log.Panic("Error with parsing file")
	// }
	// id, err := strconv.Atoi(pars.Get("username"))

	// if err != nil {
	// 	return
	// }

	var path = "./public/register.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
}

// func account(w http.ResponseWriter, r *http.Request) {
// 	templateParser, err := template.ParseFiles("./pubilc/templates/account.html")
// 	if err != nil {
// 		log.Panic("Error with parsing file")
// 	}

// 	id, err := strconv.Atoi(pars.Get("username"))

// 	if err != nil {
// 		return
// 	}

// }
