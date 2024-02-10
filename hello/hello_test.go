package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Chrish"

	got := Hello("Chrish")

	if want != want {
		t.Errorf("want %q, got %q ", want, got)
	}
}
