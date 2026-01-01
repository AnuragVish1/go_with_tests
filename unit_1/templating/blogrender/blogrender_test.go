package blogrender_test

import (
	"bytes"
	"io"
	"testing"
	"unit1/templating/blogrender"

	approvals "github.com/approvals/go-approval-tests"
)

func TestBlogRender(t *testing.T) {
	post := blogrender.Post{
		Title:       "New",
		Description: "New Post this is",
		Body:        "## New New new every thing is new here",
		Tags:        []string{"NEw", "trial"},
	}
	postren, err := blogrender.NewPostParse()

	t.Run("Converting a post in to html", func(t *testing.T) {

		buff := &bytes.Buffer{}

		if err != nil {
			t.Fatal("Got error but didnt wanted", err)
		}
		err = postren.Render(buff, post)
		if err != nil {
			t.Fatal("Got error but didnt wanted", err)
		}
		approvals.VerifyString(t, buff.String())
	})

	t.Run("Checking if the index page is rendered properly", func(t *testing.T) {
		post := []blogrender.Post{{Title: "Hello"}, {Title: "Hello2"}}
		buff := &bytes.Buffer{}

		err := postren.RenderIndex(buff, post)
		if err != nil {
			t.Fatal("Didt wanted the error but got one anyways", err)
		}

		approvals.VerifyString(t, buff.String())
	})
}

func TestRenderIndexPage(t *testing.T) {

}

func BenchmarkParse(b *testing.B) {
	post := blogrender.Post{
		Title:       "New",
		Description: "New Post this is",
		Body:        "New New new every thing is new here",
		Tags:        []string{"NEw", "trial"},
	}

	postren, err := blogrender.NewPostParse()
	if err != nil {
		b.Fatal("Got error but didnt wanted", err)
	}
	for b.Loop() {
		postren.Render(io.Discard, post)
	}

}
