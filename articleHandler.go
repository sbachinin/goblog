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

	artcl := ExtractArticleData(string(b))

	_ = artcl

	// if title == nil ||
	// 	len(title) < 2 ||
	// 	subtitle == nil ||
	// 	len(subtitle) < 2 {
	// 	continue
	// }

	// article_tpl.Execute(w, articles)
}
