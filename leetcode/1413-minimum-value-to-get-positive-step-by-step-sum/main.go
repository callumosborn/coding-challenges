package main

// https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum

// Given an array of integers nums, you start with an initial positive value startValue.
//
// In each iteration, you calculate the step by step sum of startValue plus elements
// in nums (from left to right).
//
// Return the minimum positive value of startValue such that the step by step sum
// is never less than 1.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func minStartValue(nums []int) int {
	curr := 0
	minVal := 0

	for _, v := range nums {
		curr += v
		minVal = min(minVal, curr)
	}

	return -minVal + 1
}
