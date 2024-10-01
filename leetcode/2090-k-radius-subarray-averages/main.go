package main

// https://leetcode.com/problems/k-radius-subarray-averages

// You are given a 0-indexed array nums of n integers, and an integer k.
//
// The k-radius average for a subarray of nums centered at some index i
// with the radius k is the average of all elements in nums between the
// indices i - k and i + k (inclusive). If there are less than k elements
// before or after the index i, then the k-radius average is -1.
//
// Build and return an array avgs of length n where avgs[i] is the k-radius
// average for the subarray centered at index i.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func getAverages(nums []int, k int) []int {
	avgs := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		avgs[i] = -1
	}

	if 2*k > len(nums) {
		return avgs
	}

	left := 0
	right := 0
	total := 0

	for right < 2*k {
		total += nums[right]
		right++
	}

	for right < len(nums) {
		total += nums[right]
		avgs[right-k] = total / (right - left + 1)
		right++
		total -= nums[left]
		left++
	}

	return avgs
}
