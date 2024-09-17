package main

// https://leetcode.com/problems/reverse-vowels-of-a-string/

// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u',
// and they can appear in both lower and upper cases, more than once.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'A': true,
		'a': true,
		'E': true,
		'e': true,
		'I': true,
		'i': true,
		'O': true,
		'o': true,
		'U': true,
		'u': true,
	}

	i := 0
	j := len(s) - 1

	b := []byte(s)

	for i < j {

		if vowels[s[i]] && vowels[s[j]] {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		} else {
			if !vowels[s[i]] {
				i++
			}

			if !vowels[s[j]] {
				j--
			}
		}
	}

	return string(b)
}
