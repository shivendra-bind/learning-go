package perimeter

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Ractangle", func(t *testing.T) {

		ractangle := Ractangle{5.0, 5.0}
		checkArea(t, ractangle, 25.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{5}

		want := 78.53981633974483
		checkArea(t, circle, want)
	})

}
