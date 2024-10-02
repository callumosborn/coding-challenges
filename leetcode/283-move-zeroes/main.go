package main

// https://leetcode.com/problems/move-zeroes/

// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.
//
// Time Complexity: O(N)
// Space Complexity: O(1)
func moveZeroes(nums []int) {
	i, j := 0, 1

	for j < len(nums) {

		if nums[i] == 0 && nums[j] == 0 {
			j++
		} else if nums[i] != 0 && nums[j] == 0 {
			i++
			j++
		} else if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			i++
			j++
		}
	}
}
