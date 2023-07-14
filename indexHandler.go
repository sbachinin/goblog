package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	type article struct {
		Title      string
		Subtitle   string
		Date       time.Time
		DateString string
		Url        string
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

		title_re := regexp.MustCompile(`(?m)^#{2}\s+(.+)`)
		title := title_re.FindStringSubmatch(string(b))

		subtitle_re := regexp.MustCompile(`(?m)^#{3}\s+(.+)`)
		subtitle := subtitle_re.FindStringSubmatch(string(b))

		var dateString string
		date_re := regexp.MustCompile(`^-{4}\s+(.+)`)
		date_match := date_re.FindStringSubmatch(string(b))
		if date_match != nil || len(date_match) >= 2 {
			dateString = date_match[1]
		}
		date, err := time.Parse("Jan 2 15:04:05 MST 2006", dateString)
		if err != nil {
			fmt.Println("Error parsing date:", err)
		}

		if title == nil ||
			len(title) < 2 ||
			subtitle == nil ||
			len(subtitle) < 2 {
			continue
		}

		artcl := article{
			Title:      title[1],
			Subtitle:   subtitle[1],
			Date:       date,
			DateString: date.Format("2 Jan 2006"),
			Url:        filePath,
		}
		articles = append(articles, artcl)
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[j].Date.Before(articles[i].Date)
	})

	index_tpl.Execute(w, articles)
}
