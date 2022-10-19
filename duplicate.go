package main

func ContainDuplicate(nums []int) bool {
	if nums == nil {
		return true
	}

	m := make(map[int]bool)

	for _, v := range nums {
		if m[v] {
			return true
		}

	}
	return false
}
