package internal

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var article_tpl = template.Must(template.ParseFiles("templates/article.gohtml"))

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("." + r.URL.String())
	if err != nil {
		fmt.Print(err)
	}

	// articleData := ExtractArticleData(string(b))

	// md := []byte(articleData.Content)
	// wrapper := "<div class='article-wrapper'>%v</div>"
	// fmt.Fprintf(w, wrapper, string(mdToHTML(b)))
	article_tpl.Execute(w, string(MarkdownToHTML(b)))
}
