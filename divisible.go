package main

func avaregeValue(nums []int) int {

	var sum, result int
	var divisible []int
	for _, num := range nums {
		if num%2 == 0 && num%3 == 0 {
			sum += num
			divisible = append(divisible, num)
		}
		result = sum / len(divisible)

	}

	return result

}
