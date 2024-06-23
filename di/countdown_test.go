package di

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}
func TestCountdown(t *testing.T) {
	buf := &bytes.Buffer{}
	sleeper := SpySleeper{}
	Countdown(buf, &sleeper)

	got := buf.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if sleeper.Calls != 3 {
		t.Errorf("not enough call to sleeper, want 3 got %d", sleeper.Calls)
	}
}
