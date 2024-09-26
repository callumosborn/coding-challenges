package main

import "testing"

func TestReverseStr(t *testing.T) {
	_ = reverseStr("abcdefg", 2)

	_ = reverseStr("abcd", 2)
}
