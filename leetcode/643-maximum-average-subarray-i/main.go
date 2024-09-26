package main

// https://leetcode.com/problems/maximum-average-subarray-i

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
