package main

// https://leetcode.com/problems/max-consecutive-ones-iii

// Given a binary array nums and an integer k, return the maximum number
// of consecutive 1's in the array if you can flip at most k 0's.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func longestOnes(nums []int, k int) int {
	left := 0
	curr := 0
	ans := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			curr++
		}

		for curr > k {

			if nums[left] == 0 {
				curr--
			}

			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
