package main

import "strconv"

func maximumValue(strs []string) int {
	var max int

	if len(strs) == 0 {
		return max
	}

	for _, result := range strs {
		if len(result) >= max {
			max = len(result)
		} else if _, err := strconv.Atoi(result); err == nil {
			if max < len(result) {
				max = len(result)
			}
		}
	}

	return max
}
