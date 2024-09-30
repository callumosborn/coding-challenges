package main

// https://leetcode.com/problems/max-consecutive-ones

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	curr := 0
	left := 0

	for right := 0; right < len(nums); right++ {

		if nums[right] == 0 {
			curr++
		}

		for curr > 0 {

			if nums[left] == 0 {
				curr--
			}

			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
