package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var result int
	for i, j := 1, 2; j <= n; i, j = i+1, j+1 {
		result = i + j
	}
	return result
}
