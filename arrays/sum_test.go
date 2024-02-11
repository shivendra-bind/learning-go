package arrays

import (
	"reflect"
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

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want  %v", got, want)
	}
}

func TestSumAllAppend(t *testing.T) {
	got := SumAllAppend([]int{3, 3, 3}, []int{0, 9})
	want := []int{9, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want  %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want  %v", got, want)
		}
	}
	t.Run("total of tails of all slices", func(t *testing.T) {
		got := SumAllTails([]int{3, 3, 3}, []int{0, 9})
		want := []int{6, 9}

		checkSum(t, got, want)
	})

	t.Run("total of tails of all slices and should be 0 if slice is empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}
