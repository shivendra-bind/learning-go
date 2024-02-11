package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 3, 3}, []int{0, 9})
	want := []int{9, 9}

	if got != want {
		t.Errorf("got %v want  %v", got, want)
	}
}
