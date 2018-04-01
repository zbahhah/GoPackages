// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseReverse(s string) string {
	return string(Reverse(Reverse(s)))
}

func TribleReverse(s string) string {
	return string(Reverse(Reverse(Reverse(s))))
}
