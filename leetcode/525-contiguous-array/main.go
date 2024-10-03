package main

// https://leetcode.com/problems/contiguous-array

// Given a binary array nums, return the maximum length of a contiguous subarray
// with an equal number of 0 and 1.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func findMaxLength(nums []int) int {
	counts := make(map[int]int)

	counts[0] = -1

	maxLen, count := 0, 0

	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			count++
		} else {
			count--
		}

		if _, ok := counts[count]; ok {
			maxLen = max(maxLen, i-counts[count])
		} else {
			counts[count] = i
		}
	}

	return maxLen
}
