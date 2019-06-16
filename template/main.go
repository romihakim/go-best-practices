package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type website struct {
	Title       string
	Keyword     string
	Description string
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "./login", 301)
	})
	http.HandleFunc("/login", login)
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/about", about)

	http.ListenAndServe("localhost:8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/login.html"))

	web := website{
		Title: "KodeSain | Login",
	}

	tmpl.Execute(w, &web)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("html/layout.html", "html/dashboard.html")

	web := website{
		Title: "KodeSain | Dashboard",
	}

	tmpl.Execute(w, &web)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("html/layout.html", "html/about.html")

	str := []byte(`[{"host": "0.0.0.0", "job": "AAA", "status": "up", "date": "10/09/2018"}, {"host": "0.0.0.0", "job": "BBB", "status": "up", "date": "11/09/2018"}, {"host": "0.0.0.0", "job": "CCC", "status": "up", "date": "12/09/2018"}]`)
	info := make([]map[string]interface{}, 0)
	json.Unmarshal(str, &info)

	data := struct {
		Title string
		Name  string
		Email string
		Info  []map[string]interface{}
	}{
		"KodeSain | About",
		"ROMI",
		"romi@mail.com",
		info,
	}

	tmpl.Execute(w, &data)
}
