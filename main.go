package main

import (
	"html/template"
	"net/http"
	"os"
)

var index_tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/articles/", ArticleHandler)
	http.ListenAndServe(":"+port, mux)
}
