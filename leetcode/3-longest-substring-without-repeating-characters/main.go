package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string s, find the length of the longest substring without repeating characters.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func lengthOfLongestSubstring(s string) int {
	counts := make(map[byte]int)

	left, ans := 0, 0

	for right := 0; right < len(s); right++ {
		counts[s[right]]++

		for counts[s[right]] > 1 {
			counts[s[left]]--
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
