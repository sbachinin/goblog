package main

import (
	"fmt"
	"net/http"
	"os"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("." + r.URL.String())
	if err != nil {
		fmt.Print(err)
	}

	w.Write(b)
}
