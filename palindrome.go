package main

func PalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	var reverse = 0

	temp := x

	for temp > 0 {
		reverse = reverse*10 + temp%10
		temp = temp / 10
	}
	return x == reverse
}
