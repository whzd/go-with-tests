package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}
	got := Sum(numbers)
	want := 6
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 9, 12})
	want := []int{6, 21}
	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("sum some slices tails", func(t *testing.T) {
		got := SumAllTails([]int{0, 7, 9}, []int{3, 14, 17})
		want := []int{16, 31}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 14, 17})
		want := []int{0, 31}
		checkSums(t, got, want)
	})
}
