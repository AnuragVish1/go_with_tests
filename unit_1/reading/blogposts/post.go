package blogposts

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

// type FailingFileSystem struct {
// }

// func (ffs *FailingFileSystem) Open(name string) (fs.File, error) {
// 	return nil, errors.New("Got the error bro")
// }

const (
	titleStart       = "Title: "
	descriptionStart = "Description: "
	tagsStart        = "Tags: "
	bodyStart        = "---"
)

func NewPostFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}
	posts := []Post{}

	for _, file := range dir {
		post, err := getPost(filesystem, file.Name())
		if err != nil {
			return nil, errors.New("This i cant do")
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, file string) (Post, error) {
	postFile, err := fileSystem.Open(file)

	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	return makeNewPost(postFile)
}

func makeNewPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readMetaLine := func(seprator string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), seprator)
	}

	titleText := readMetaLine(titleStart)

	descriptionText := readMetaLine(descriptionStart)

	tags := strings.Split(readMetaLine(tagsStart), ",")

	return Post{
			Title:       titleText,
			Description: descriptionText,
			Tags:        tags,
			Body:        readBody(scanner)},
		nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buff := &bytes.Buffer{}
	if scanner.Text() == bodyStart {

		for scanner.Scan() {
			fmt.Fprintln(buff, scanner.Text())
		}
	}
	body := strings.TrimSuffix(buff.String(), "\n")
	return body

}
