package main

func MaxSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
