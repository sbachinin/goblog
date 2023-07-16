package internal

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var article_tpl = template.Must(template.ParseFiles("templates/article.gohtml"))

func fixDate(articleString string, dateString string) string {
	/* title_start_index */ t_s_i := strings.Index(articleString, "<h2")
	/* content_start_index */ c_s_i := strings.Index(articleString, "<h4")
	dateString = "<p class='article-date'>" + dateString + "</p>"
	if c_s_i != -1 {
		return articleString[t_s_i:c_s_i] + dateString + articleString[c_s_i:]
	} else {
		return articleString[t_s_i:] + dateString
	}
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("." + r.URL.String())
	if err != nil {
		fmt.Print(err)
	}

	articleData := ExtractArticleData(string(b))

	article_html := string(MarkdownToHTML(b))
	article_html = fixDate(article_html, articleData.DateString)

	article_tpl.Execute(w, article_html)
}
