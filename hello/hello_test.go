package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		want := "Hello, Chrish"

		got := Hello("Chrish")

		if want != want {
			t.Errorf("want %q, got %q ", want, got)
		}
	})
	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("")

		if want != got {
			t.Errorf("want %q, got %q ", want, got)
		}
	})
}
