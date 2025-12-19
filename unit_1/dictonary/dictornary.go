package dictonary

type Dictonary map[string]string

const (
	ErrorMessage       = DictonaryErr("no word found")
	ErrorAlreadyExists = DictonaryErr("word already exists")
)

type DictonaryErr string

func (e DictonaryErr) Error() string {
	return string(e)
}

func (d Dictonary) Search(word string) (string, error) {
	defination, ok := d[word]
	if !ok {
		return "", ErrorMessage
	}
	return defination, nil

}

func (d Dictonary) Add(key string, value string) error {
	_, ok := d[key]
	if !ok {
		d[key] = value
		return nil
	}

	return ErrorAlreadyExists
}

func (d Dictonary) Replace(word string, defination string) error {
	_, err := d[word]
	if !err {
		return ErrorMessage
	}
	d[word] = defination
	return nil
}

func (d Dictonary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorMessage:
		return ErrorMessage
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}

func main() {

}

func Search(dict map[string]string, word string) string {
	return dict[word]
}
