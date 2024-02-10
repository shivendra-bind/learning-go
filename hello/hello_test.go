package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, World"

	got := Hello()

	if want != want {
		t.Errorf("want %q, got %q ", want, got)
	}
}
