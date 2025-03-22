package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("calculate sum of each array", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{0, 9, 12})
		want := []int{6, 21}
		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
