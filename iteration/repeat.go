package iteration

import "strings"

// Repeat repeats the input string 'number' number of times
func Repeat(chars string, number int) string {
	var repeated strings.Builder
	for range number {
		repeated.WriteString(chars)
	}
	return repeated.String()
}
