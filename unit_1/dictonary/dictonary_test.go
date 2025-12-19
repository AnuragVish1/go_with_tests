package dictonary

import (
	"testing"
)

func TestDictonary(t *testing.T) {
	t.Run("Dictonary working correctly", func(t *testing.T) {
		dictoary := Dictonary{"test": "hellos"}
		got, _ := dictoary.Search("test")
		want := "hellos"

		validate(got, want, t)
	})

	t.Run("Word not found", func(t *testing.T) {
		dict := Dictonary{"test": "here"}
		_, err := dict.Search("yeah")
		validateError(err, t, ErrorMessage)
		validate(err.Error(), ErrorMessage.Error(), t)

	})

	t.Run("word already exists", func(t *testing.T) {
		word := "hello"
		defination := "hi"

		dict := Dictonary{word: defination}

		err := dict.Add(word, "hey")

		validateError(err, t, ErrorAlreadyExists)

	})

}

func TestAddWord(t *testing.T) {
	t.Run("Add word to Dictonary", func(t *testing.T) {
		dict := Dictonary{"test": "helo"}

		word := "yes"
		defination := "sir"

		dict.Add(word, defination)
		validateDefination(dict, defination, word, t)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	dict := Dictonary{word: "Hello"}
	replacedDefination := "hi"
	t.Run("Updating the defination", func(t *testing.T) {
		dict.Replace(word, replacedDefination)
		updatedDefination, _ := dict.Search(word)
		validate(updatedDefination, replacedDefination, t)
	})

	t.Run("Updating the word that doesnt exists", func(t *testing.T) {
		err := dict.Replace("he", replacedDefination)
		validateError(err, t, ErrorMessage)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictonary{word: "Hello"}
	dict.Delete(word)
	_, err := dict.Search(word)
	validateError(err, t, ErrorMessage)
}

func validate(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func validateError(got error, t *testing.T, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Expected Error but got none")
	}

	if got.Error() != want.Error() {
		t.Errorf("Got %q want %q", got, want.Error())
	}
}

func validateDefination(dict Dictonary, defination string, word string, t *testing.T) {

	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("Error, should add word")
	}

	validate(got, defination, t)
}
