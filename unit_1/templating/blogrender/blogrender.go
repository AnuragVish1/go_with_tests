package blogrender

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func main() {

}

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func (p Post) ToHTML() template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	par := parser.NewWithExtensions(extensions)
	doc := par.Parse([]byte(p.Body))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank

	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	fmt.Print(string(markdown.Render(doc, renderer)))
	return template.HTML(string(markdown.Render(doc, renderer)))
}

var (
	//go:embed "templates/*"
	htmlSkeleton embed.FS
)

type PostRender struct {
	templ *template.Template
}

func NewPostParse() (*PostRender, error) {
	tem, err := template.ParseFS(htmlSkeleton, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &PostRender{templ: tem}, nil
}

func (p *PostRender) Render(writer io.Writer, post Post) error {
	if err := p.templ.ExecuteTemplate(writer, "blog.gohtml", post); err != nil {
		return err
	}
	return nil
}
