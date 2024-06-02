package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is just a test"}
		got, _ := dictionary.Search("test")
		want := "This is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		dictionary := Dictionary{"test": "This is just a test"}
		_, err := dictionary.Search("test1")
		assertError(t, err, ErrNotFound)

	})
}

func TestAddWord(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dict := Dictionary{}
		word := "test"
		def := "This is a word test"

		dict.Add(word, def)
		assertDefination(t, dict, word, def)
	})

	t.Run("existing word", func(t *testing.T) {

		dict := Dictionary{}
		word := "test"
		def := "This is a word test"

		dict.Add(word, def)
		err := dict.Add(word, "override word")

		assertDefination(t, dict, word, def)
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def := "This is a word test"
		newDef := "This is a new word"

		dict.Add(word, def)
		dict.Update(word, newDef)

		assertDefination(t, dict, word, newDef)
	})
}
func assertDefination(t testing.TB, dict Dictionary, word, defination string) {
	t.Helper()

	got, _ := dict.Search(word)

	assertString(t, got, defination)

}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q , given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q , given %q", got, want, "test")
	}
}
