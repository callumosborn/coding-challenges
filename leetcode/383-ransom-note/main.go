package main

// https://leetcode.com/problems/ransom-note/

// Given two strings ransomNote and magazine,
// return true if ransomNote can be constructed by using
// the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.
//
// Time Complexity: O(N + M)
// Space Complexity: O(N)
func canConstruct(ransomNote string, magazine string) bool {
	counts := make(map[rune]int)

	for _, v := range magazine {
		counts[v]++
	}

	for _, v := range ransomNote {

		if counts[v] == 0 {
			return false
		}

		counts[v]--
	}

	return true
}
