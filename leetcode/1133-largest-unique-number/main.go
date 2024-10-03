package main

// https://leetcode.com/problems/largest-unique-number/

// Given an integer array nums, return the largest integer that only occurs once.
// If no integer occurs once, return -1.
//
// Time Complexity: O(N)
// Space Compleixty: O(N)
func largestUniqueNumber(nums []int) int {
	counts := make(map[int]int)

	for _, v := range nums {

		cv, ok := counts[v]

		if !ok {
			counts[v] = 1
		} else {
			counts[v] = cv + 1
		}
	}

	ans := -1

	for k, v := range counts {

		if v == 1 {
			ans = max(ans, k)
		}
	}

	return ans
}
