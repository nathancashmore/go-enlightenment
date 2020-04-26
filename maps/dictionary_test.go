package maps

import (
	"fmt"
	"testing"
)

func ExampleDictionary_Search() {
	dictionary := Dictionary{"term": "description"}
	result, _ := dictionary.Search("term")

	fmt.Println(result)
	// Output: description
}

func BenchmarkDictionary_Search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dictionary := Dictionary{"term": "description"}
		dictionary.Search("term")
	}
}

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(got, want, t)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("random")

		assertError(ErrNotFound, err, t)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add should add", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("word", "a definition")

		want := "a definition"
		got, err := dictionary.Search("word")

		assertStrings(got, want, t)

		if err != nil {
			t.Fatalf("Received an error when there should have been none - %s", err)
		}
	})

	t.Run("add should not overwrite", func(t *testing.T) {
		dictionary := Dictionary{"one": "the first one"}

		err := dictionary.Add("one", "the second one")

		assertError(err, ErrAddAlreadyExists, t)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update entry", func(t *testing.T) {
		dictionary := Dictionary{"update": "update me please"}

		dictionary.Update("update", "has been updated")

		want := "has been updated"
		got, _ := dictionary.Search("update")

		assertStrings(got, want, t)
	})

	t.Run("cannot update if it does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("update", "has been updated")
		want := ErrUpdateDoesNotExist
		assertError(err, want, t)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertError(err, ErrNotFound, t)
}

func assertError(got error, want error, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("Wrong error : got - %s, want - %s", got, want)
	}

	if got == nil {
		t.Errorf("Expected an error but did not get one")
	}
}

func assertStrings(got string, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
