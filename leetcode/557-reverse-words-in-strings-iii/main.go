package main

import "strings"

// https://leetcode.com/problems/reverse-words-in-a-string-iii/

// Given a string s, reverse the order of characters in each word within a sentence
// while still preserving whitespace and initial word order.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func reverseWords(s string) string {
	w := strings.Split(s, " ")

	for h := 0; h < len(w); h++ {

		b := []byte(w[h])

		i := 0
		j := len(b) - 1

		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}

		w[h] = string(b)
	}

	return strings.Join(w, " ")
}
