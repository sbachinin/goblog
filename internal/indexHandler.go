package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
)

var index_tpl = template.Must(template.ParseFiles("templates/index.gohtml"))

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

	for _, entry := range entries {

		if !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filePath := "./articles/dev/" + entry.Name()

		b, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Print("Error when reading entry in articles/dev folder: ", err)
			continue
		}

		filePath = strings.TrimSuffix(filePath, ".md")

		artcl := article{
			Data: ExtractArticleData(string(b)),
			Url:  filePath,
		}

		if len(artcl.Data.Title) > 0 && artcl.Data.HaveContent {
			articles = append(articles, artcl)
		}
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[j].Data.Date.Before(articles[i].Data.Date)
	})

	index_tpl.Execute(w, articles)
}
