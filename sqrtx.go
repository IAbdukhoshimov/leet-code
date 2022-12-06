package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var result int
	for i := 1; i <= x; i++ {
		if i*i <= x {
			result = i
		} else {
			break
		}
	}
	return result
}
