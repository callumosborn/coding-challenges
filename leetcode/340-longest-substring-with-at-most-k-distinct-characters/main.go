package main

// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/

// Given a string s and an integer k, return the length of the longest  substring of s
// that contains at most k distinct characters.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	counts := make(map[byte]int)

	left, ans := 0, 0

	for right := 0; right < len(s); right++ {
		counts[s[right]]++

		for len(counts) > k {
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
