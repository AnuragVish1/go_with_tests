package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/AnuragVish1/blogposts"
)

const (
	firstFileText = `Title: Thisis the hett
Description: This is the post about AI
Tags: AI,ML
---
this
is
the
body of t
d `
	secondFileText = `Title: This is the new blog
Description: This is the blog for my go journey
Tags: Yo,Yoo
---
this
is
second
post
tru`
)

func TestBlogFiles(t *testing.T) {
	t.Run("Checking if we getting correct data", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"hey.md":   {Data: []byte(firstFileText)},
			"hello.md": {Data: []byte(secondFileText)},
		}

		posts, err := blogposts.NewPostFromFS(fileSystem)

		if err != nil {
			t.Fatal("Shouldn't have gotten error but got one anyways", err)
		}

		got := posts[1]
		want := blogposts.Post{
			Title:       "Thisis the hett",
			Description: "This is the post about AI",
			Tags:        []string{"AI", "ML"},
			Body: `this
is
the
body of t
d `}

		validatePosts(t, got, want)

	})
}

func validatePosts(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}
