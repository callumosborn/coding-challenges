package main

// https://leetcode.com/problems/reverse-only-letters

// Given a string s, reverse the string according to the following rules:
//
// All the characters that are not English letters remain in the same position.
// All the English letters (lowercase or uppercase) should be reversed.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func reverseOnlyLetters(s string) string {
	b := []byte(s)

	i, j := 0, len(b)-1

	for i < j {

		for ((b[i] >= 33 && b[i] < 65) || (b[i] >= 91 && b[i] < 97)) && i < len(b)-1 {
			i++
		}

		if i == len(b)-1 {
			break
		}

		for (b[j] >= 33 && b[j] < 65) || (b[j] >= 91 && b[j] < 97) {
			j--
		}

		if i >= j {
			break
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}
