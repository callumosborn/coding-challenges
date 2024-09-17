package main

// https://leetcode.com/problems/reverse-string/

// Write a function that reverses a string.
// The input string is given as an array of characters s.
//
// You must do this by modifying the input array in-place with O(1) extra memory.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]

		i++
		j--
	}
}
