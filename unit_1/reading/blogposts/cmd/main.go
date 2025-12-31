package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/AnuragVish1/blogposts"
)

func main() {
	posts, err := blogposts.NewPostFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}

	for i, post := range posts {
		log.Printf("Following is the post data for post %d", i+1)
		b, err := json.MarshalIndent(post, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(b))
	}
}
