package main

import "testing"

func TestReverseString(t *testing.T) {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})

	reverseString([]byte{'H', 'a', 'n', 'n', 'a', 'h'})
}
