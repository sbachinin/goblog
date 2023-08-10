package internal

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var article_tpl = template.Must(template.ParseFiles("templates/article.gohtml"))

func insertNiceDate(articleString string, dateString string) string {
	if dateString == "" {
		return ""
	}

	dateString = "<p class='article-date'>" + dateString + "</p>"

	// insert the date after subtitle or, if there is
	date_index := -1

	subtitle_close_index := strings.Index(articleString, "</h2>")
	if subtitle_close_index != -1 {
		date_index = subtitle_close_index + 5
	} else {
		title_close_index := strings.Index(articleString, "</h1>")
		if title_close_index != -1 {
			date_index = title_close_index + 5
		}
	}

	// Insert date before content if there is content
	// otherwise just append date
	if date_index != -1 {
		articleString = articleString[:date_index] + dateString + articleString[date_index:]
	}

	return articleString
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := os.ReadFile("." + filepath.Clean(r.URL.Path) + ".md")
	if err != nil {
		fmt.Print(err)
	}

	articleData := ExtractArticleData(string(bytes))

	articleHTML := string(MarkdownToHTML(bytes))

	articleIsFinished := false

	titleIndex := strings.Index(articleHTML, "<h1")
	if titleIndex != -1 {
		// Remove a line with raw date
		articleHTML = articleHTML[titleIndex:]

		if strings.Contains(articleHTML, "</p>") {
			articleIsFinished = true
		}
	}

	if !articleIsFinished {
		article_tpl.Execute(w, "<p class='nothing'>This article is not written yet (</p>")
		return
	}

	articleHTML = insertNiceDate(articleHTML, articleData.DateString)

	data := struct {
		Title       string
		ArticleHTML string
	}{
		Title:       articleData.Title,
		ArticleHTML: articleHTML,
	}

	article_tpl.Execute(w, data)
}
