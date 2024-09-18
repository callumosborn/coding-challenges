package main

import (
	"bytes"
)

// https://leetcode.com/problems/valid-palindrome

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
// and removing all non-alphanumeric characters, it reads the same forward and backward.
//
// Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func isPalidrone(s string) bool {
	i := 0
	j := len(s) - 1

	b := bytes.ToLower([]byte(s))

	for i < j {

		for !(b[i] >= 48 && b[i] <= 57) && !(b[i] >= 97 && b[i] <= 122) && i < j {
			i++
		}

		for !(b[j] >= 48 && b[j] <= 57) && !(b[j] >= 97 && b[j] <= 122) && i < j {
			j--
		}

		if b[i] != b[j] {
			return false
		}

		i++
		j--
	}

	return true
}
