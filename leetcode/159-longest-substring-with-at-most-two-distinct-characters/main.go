package main

// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/

// Given a string s, return the length of the longest substring that contains
// at most two distinct characters.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func lengthOfLongestSubstringTwoDistinct(s string) int {
	counts := make(map[byte]int)

	left, ans := 0, 0

	for right := 0; right < len(s); right++ {
		counts[s[right]]++

		for len(counts) > 2 {
			counts[s[left]]--

			if counts[s[left]] == 0 {
				delete(counts, s[left])
			}

			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
