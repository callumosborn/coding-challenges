package main

// https://leetcode.com/problems/check-if-the-sentence-is-pangram

// A pangram is a sentence where every letter of the English alphabet appears at least once.
// Given a string sentence containing only lowercase English letters,
// return true if sentence is a pangram, or false otherwise.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func checkIfPangram(sentence string) bool {
	hash := map[rune]bool{}

	for _, c := range sentence {
		hash[c] = true
	}

	return len(hash) == 26
}
