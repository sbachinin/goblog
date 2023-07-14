package main

import (
	"fmt"
	"regexp"
	"time"
)

type articleData struct {
	Title      string
	Subtitle   string
	Date       time.Time
	DateString string
}

func ExtractArticleData(articleText string) articleData {
	title_re := regexp.MustCompile(`(?m)^#{2}\s+(.+)`)
	title := title_re.FindStringSubmatch(articleText)

	subtitle_re := regexp.MustCompile(`(?m)^#{3}\s+(.+)`)
	subtitle := subtitle_re.FindStringSubmatch(articleText)

	var dateString string
	date_re := regexp.MustCompile(`^-{4}\s+(.+)`)
	date_match := date_re.FindStringSubmatch(articleText)
	if date_match != nil || len(date_match) >= 2 {
		dateString = date_match[1]
	}
	date, err := time.Parse("Jan 2 15:04:05 MST 2006", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	return articleData{
		Title:      title[1],
		Subtitle:   subtitle[1],
		Date:       date,
		DateString: date.Format("2 Jan 2006"),
	}
}
