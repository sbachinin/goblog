package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	entries, err := os.ReadDir("./articles/dev")
	if err != nil {
		log.Fatal(err)
	}

	type article struct {
		Title    string
		Subtitle string
		Url      string
	}

	articles := []article{}

	for _, e := range entries {
		filePath := "./articles/dev/" + e.Name()
		b, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Print(err)
		}

		title_re := regexp.MustCompile(`(?m)^#{2}\s+(.+)`)
		title := title_re.FindStringSubmatch(string(b))

		subtitle_re := regexp.MustCompile(`(?m)^#{3}\s+(.+)`)
		subtitle := subtitle_re.FindStringSubmatch(string(b))

		if title == nil ||
			len(title) < 2 ||
			subtitle == nil ||
			len(subtitle) < 2 {
			continue
		}

		artcl := article{
			Title:    title[1],
			Subtitle: subtitle[1],
			Url:      filePath,
		}
		articles = append(articles, artcl)
	}

	tpl.Execute(w, articles)
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
	http.ListenAndServe(":"+port, mux)
}
