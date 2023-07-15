package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var article_tpl = template.Must(template.ParseFiles("article.gohtml"))

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("." + r.URL.String())
	if err != nil {
		fmt.Print(err)
	}

	articleData := ExtractArticleData(string(b))

	article_tpl.Execute(w, articleData)
}
