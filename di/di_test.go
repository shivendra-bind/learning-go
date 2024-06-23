package di

import (
	"bytes"
	"os"
	"testing"
)

func TestGreetWithBytes(t *testing.T) {
	buf := bytes.Buffer{}

	Greet(&buf, "Chrish")

	got := buf.String()
	want := "Hello, Chrish"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreetWithIO(t *testing.T) {
	Greet(os.Stdout, "Chrish")
}
