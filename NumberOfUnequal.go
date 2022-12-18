package main

func unequalTriplets(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] != nums[j] && nums[j] != nums[k] && nums[i] != nums[k] {
					count++
				}
			}
		}
	}
	return count
}
