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
	if len(dateString) > 0 {
		dateString = "<p class='article-date'>" + dateString + "</p>"
	}

	title_index := strings.Index(articleString, "<h2 ")

	// Remove all before title if there is title
	// (What's before is probably a line with "----" & raw date)
	if title_index != -1 {
		articleString = articleString[title_index:]
	}

	content_index := strings.Index(articleString, "<h4 ")

	// Insert date before content if there is content
	// otherwise just append date
	if content_index != -1 {
		articleString = articleString[:content_index] + dateString + articleString[content_index:]
	} else {
		articleString = articleString + dateString
	}

	return articleString
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("." + r.URL.String() + ".md")
	if err != nil {
		fmt.Print(err)
	}

	articleData := ExtractArticleData(string(b))
	if len(articleData.Title) == 0 || len(articleData.Content) == 0 {
		article_tpl.Execute(w, "<p class='nothing'>This article is not written yet (</p>")
		return
	}

	// else: there is title and content, get html

	article_html := string(MarkdownToHTML(b))

	article_html = fixDate(article_html, articleData.DateString)

	article_tpl.Execute(w, article_html)
}
