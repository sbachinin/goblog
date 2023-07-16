package internal

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

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
	article_tpl.Execute(w, string(mdToHTML(b)))
}
