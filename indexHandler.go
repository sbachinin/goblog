package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	type article struct {
		Data articleData
		Url  string
	}

	articles := []article{}

	entries, err := os.ReadDir("./articles/dev")
	if err != nil {
		log.Fatal(err)
	}

	for i := len(entries) - 1; i >= 0; i-- {
		filePath := "./articles/dev/" + entries[i].Name()
		b, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Print(err)
		}

		artcl := article{
			Data: ExtractArticleData(string(b)),
			Url:  filePath,
		}

		if len(artcl.Data.Title) > 0 {
			articles = append(articles, artcl)
		}
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[j].Data.Date.Before(articles[i].Data.Date)
	})

	index_tpl.Execute(w, articles)
}
