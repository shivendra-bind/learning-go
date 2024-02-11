package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat charecter 5 times", func(t *testing.T) {
		got := Repeat("q")

		expected := "qqqqq"
		if got != expected {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a")
	}

}
