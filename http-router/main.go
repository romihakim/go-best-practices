package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/app", page)
	router.GET("/app/:page", page)

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	router.ServeFiles("/content/*filepath", http.Dir("content"))

	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "./app/index", 301)
}

func page(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	page := "ïndex"

	if p.ByName("page") != "" {
		page = p.ByName("page")
	}

	app := map[string]interface{}{
		"Title":   strings.ToUpper(page),
		"App":     "http://localhost:8080/app/",
		"Assets":  "http://localhost:8080/assets/",
		"Content": "http://localhost:8080/content/",
	}

	tmpl, err := template.ParseFiles(page + ".html")

	if err != nil {
		w.WriteHeader(404)
		return
	}

	tmpl.Execute(w, &app)
}
