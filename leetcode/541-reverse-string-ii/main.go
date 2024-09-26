package main

// https://leetcode.com/problems/reverse-string-ii/

// Given a string s and an integer k,
// reverse the first k characters for every 2k characters
// counting from the start of the string.
//
// If there are fewer than k characters left, reverse all of them.
// If there are less than 2k but greater than or equal to k characters,
// then reverse the first k characters and leave the other as original.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func reverseStr(s string, k int) string {
	b := []byte(s)

	for h := 0; h < len(s); h = h + 2*k {
		i := h
		j := min(h+k-1, len(s)-1)

		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}

	return string(b)
}
