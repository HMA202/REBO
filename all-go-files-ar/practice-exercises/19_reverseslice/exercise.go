package main

func ReverseSlice(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = nums[len(nums)-1-i]
	}
	return result
}
