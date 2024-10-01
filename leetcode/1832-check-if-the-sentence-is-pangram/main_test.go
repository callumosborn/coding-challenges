package main

import "testing"

func TestCheckIfPangram(t *testing.T) {
	_ = checkIfPangram("thequickbrownfoxjumpsoverthelazydog")

	_ = checkIfPangram("leetcode")
}
