package main

import "testing"

func TestReverseOnlyLetters(t *testing.T) {
	_ = reverseOnlyLetters("ab-cd")

	_ = reverseOnlyLetters("a-bC-dEf-ghIj")

	_ = reverseOnlyLetters("tNH95P=TV")

	_ = reverseOnlyLetters("7_28]")

	_ = reverseOnlyLetters("(;,9=/'*")
}
