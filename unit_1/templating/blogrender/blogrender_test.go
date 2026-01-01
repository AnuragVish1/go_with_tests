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

	t.Run("Converting a post in to html", func(t *testing.T) {

		buff := &bytes.Buffer{}
		postren, err := blogrender.NewPostParse()
		if err != nil {
			t.Fatal("Got error but didnt wanted", err)
		}
		err = postren.Render(buff, post)
		if err != nil {
			t.Fatal("Got error but didnt wanted", err)
		}
		approvals.VerifyString(t, buff.String())
	})
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
