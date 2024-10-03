package main

// https://leetcode.com/problems/maximum-number-of-balloons

// Given a string text, you want to use the characters of text
// to form as many instances of the word "balloon" as possible.
//
// Time Complexity: O(N + M)
// Space Complexity: O(P)
func maxNumberOfBalloons(text string) int {
	counts := make(map[byte]int)

	for i := 0; i < len(text); i++ {
		l := text[i]

		if l == 'b' || l == 'a' || l == 'l' || l == 'o' || l == 'n' {
			counts[l]++
		}
	}

	ans := 0

	if counts['b'] == 0 {
		return ans
	}

	for i := 0; i < counts['b']; i++ {

		if counts['a']/1 > 0 &&
			counts['l']/2 > 0 &&
			counts['o']/2 > 0 &&
			counts['n']/1 > 0 {
			ans++
		}

		counts['a']--
		counts['l'] -= 2
		counts['o'] -= 2
		counts['n']--
	}

	return ans
}
