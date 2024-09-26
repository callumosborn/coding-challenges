package main

// https://leetcode.com/problems/maximum-average-subarray-i

// You are given an integer array nums consisting of n elements, and an integer k.
//
// Find a contiguous subarray whose length is equal to k that has the maximum average value
// and return this value. Any answer with a calculation error less than 10-5 will be accepted.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func findMaxAverage(nums []int, k int) float64 {
	var curr float64

	for i := 0; i < k; i++ {
		curr += float64(nums[i])
	}

	ans := curr / float64(k)

	for i := k; i < len(nums); i++ {
		curr += float64(nums[i])
		curr -= float64(nums[i-k])

		ans = max(ans, curr/float64(k))
	}

	return ans
}
