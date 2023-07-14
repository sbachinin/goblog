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
	title := ""
	title_re := regexp.MustCompile(`(?m)^#{2}\s+(.+)`)
	title_match := title_re.FindStringSubmatch(articleText)
	if title_match != nil && len(title_match) >= 2 {
		title = title_match[1]
	}

	subtitle := ""
	subtitle_re := regexp.MustCompile(`(?m)^#{3}\s+(.+)`)
	subtitle_match := subtitle_re.FindStringSubmatch(articleText)
	if subtitle_match != nil && len(subtitle_match) >= 2 {
		subtitle = subtitle_match[1]
	}

	dateToRender := ""
	var date time.Time
	date_re := regexp.MustCompile(`^-{4}\s+(.+)`)
	date_match := date_re.FindStringSubmatch(articleText)
	if date_match != nil || len(date_match) >= 2 {
		var err error
		date, err = time.Parse("Jan 2 15:04:05 MST 2006", date_match[1])
		if err != nil {
			fmt.Println("Error parsing date:", err)
		} else {
			dateToRender = date.Format("2 Jan 2006")
		}
	}

	aData := articleData{
		Title:      title,
		Subtitle:   subtitle,
		DateString: dateToRender,
		Date:       date,
	}

	fmt.Println("date string:", aData.DateString)

	return aData
}
