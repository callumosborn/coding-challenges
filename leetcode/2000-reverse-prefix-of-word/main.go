package main

// https://leetcode.com/problems/reverse-prefix-of-word/

// Given a 0-indexed string word and a character ch,
// reverse the segment of word that starts at index 0 and ends at the index of the first occurrence
// of ch (inclusive). If the character ch does not exist in word, do nothing.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func reversePrefix(word string, ch byte) string {
	b := []byte(word)

	i := 0

	for j := 0; j < len(b); j++ {

		if b[j] == ch {

			for i < j {
				b[i], b[j] = b[j], b[i]
				i++
				j--
			}

			break
		}
	}

	return string(b)
}
