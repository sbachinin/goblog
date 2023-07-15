package main

import (
	"goblog/internal"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", internal.IndexHandler)
	mux.HandleFunc("/articles/", internal.ArticleHandler)
	http.ListenAndServe(":"+port, mux)
}
