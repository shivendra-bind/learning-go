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
		assertErrot(t, err, ErrNotFound)

	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q , given %q", got, want, "test")
	}
}

func assertErrot(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q , given %q", got, want, "test")
	}
}
