package main

// https://leetcode.com/problems/first-letter-to-appear-twice

// Given a string s consisting of lowercase English letters,
// return the first letter to appear twice.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func repeatedCharacter(s string) byte {
	hash := map[byte]bool{}

	for i := 0; i < len(s); i++ {

		if hash[s[i]] {
			return s[i]
		} else {
			hash[s[i]] = true
		}
	}

	return byte(0)
}
