package main

import "math"

func pivotInteger(n int) int {
	var x float64

	x = math.Sqrt(float64((n*n + n) / 2))
	if math.Floor(x) == x {
		return int(x)
	}
	return -1
}
