package blogposts

import (
	"errors"
	"io/fs"
)

func main() {

}

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
