package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
