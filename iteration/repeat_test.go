package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
	t.Run("return empty string if repeat number is 0", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleRepeat() {
	res := Repeat("op", 3)
	fmt.Println(res)
	// Output: opopop
}

var result string

func BenchmarkRepeat(b *testing.B) {
	var r string
	for b.Loop() {
		r = Repeat("a", 5)
	}
	result = r
}
