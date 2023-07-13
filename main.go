package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

var index_tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	type article struct {
		Title    string
		Subtitle string
		Date     string
		Url      string
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
		date, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", dateString)
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
			Title:    title[1],
			Subtitle: subtitle[1],
			Date:     date.Format("2 Jun 2006"),
			Url:      filePath,
		}
		articles = append(articles, artcl)
	}

	index_tpl.Execute(w, articles)
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("." + r.URL.String())
	if err != nil {
		fmt.Print(err)
	}

	w.Write(b)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/articles/", articleHandler)
	http.ListenAndServe(":"+port, mux)
}
