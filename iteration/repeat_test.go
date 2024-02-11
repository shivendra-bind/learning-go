package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat charecter 5 times", func(t *testing.T) {
		got := Repeat("q", 5)

		expected := "qqqqq"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("repeat charecter given no. of times", func(t *testing.T) {
		iteration := 10

		got := Repeat("r", iteration)
		expected := "rrrrrrrrrr"

		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}

}
