package main

import (
	"html/template"
	"net/http"
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

	var path = "./public/kino.html"
	var filename = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	t, _ := template.ParseFiles(path)
	t.ExecuteTemplate(w, filename, nil)
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
