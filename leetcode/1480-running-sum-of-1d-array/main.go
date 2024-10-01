package main

// https://leetcode.com/problems/running-sum-of-1d-array/

// Given an array nums. We define a running sum of an array as
// runningSum[i] = sum(nums[0]…nums[i]).
//
// Return the running sum of nums.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func runningSum(nums []int) []int {
	prefix := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		prefix = append(prefix, nums[i]+prefix[i-1])
	}

	return prefix
}
