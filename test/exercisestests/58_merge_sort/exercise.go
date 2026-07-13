package main

func merge(left []int, right []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		copyNums := make([]int, len(nums))
		copy(copyNums, nums)
		return copyNums
	}

	mid := len(nums) / 2

	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	return merge(left, right)
}
