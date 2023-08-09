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

	stylesFS := http.FileServer(http.Dir("styles"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", stylesFS))

	articleAssetsFS := http.FileServer(http.Dir("articles/dev/assets"))
	mux.Handle("/articles/dev/assets/", http.StripPrefix("/articles/dev/assets/", articleAssetsFS))

	mux.HandleFunc("/", internal.IndexHandler)
	mux.HandleFunc("/articles/", internal.ArticleHandler)
	http.ListenAndServe(":"+port, mux)
}
