package internal

import (
	"fmt"
	"strings"
	"time"
)

type articleData struct {
	Title       string
	Subtitle    string
	Date        time.Time
	DateString  string
	HaveContent bool
}

func ExtractArticleData(articleText string) articleData {

	title := ""
	subtitle := ""
	var date time.Time
	dateToRender := ""
	haveContent := false

	lines := strings.Split(articleText, "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "# ") {
			title = line[2:]
		} else if strings.HasPrefix(line, "## ") {
			subtitle = line[3:]
		} else if strings.HasPrefix(line, "---- ") {
			var err error
			date, err = time.Parse("Jan 2 15:04:05 MST 2006", line[5:])
			if err != nil {
				fmt.Println("Error parsing date:", err)
			} else {
				dateToRender = date.Format("2 Jan 2006")
			}
		} else if len(line) > 0 {
			// there is a line which is not title/subtitle/date
			haveContent = true
		}
	}

	aData := articleData{
		Title:       title,
		Subtitle:    subtitle,
		DateString:  dateToRender,
		Date:        date,
		HaveContent: haveContent,
	}

	return aData
}
