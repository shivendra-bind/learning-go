package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Chrish"

		got := Hello("Chrish")

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("")

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello in Spanish", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", "Spanish")

		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in French", func(t *testing.T) {
		want := "Bonjour, Ralf"
		got := Hello("Ralf", "French")

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want %q, got %q ", want, got)
	}
}
