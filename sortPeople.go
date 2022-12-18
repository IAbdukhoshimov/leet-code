package main

func sortPeople(names []string, heights []int) []string {
	var temp string
	var temp2 int

	for i := 0; i < len(heights); i++ {
		for j := i + 1; j < len(heights); j++ {
			if heights[i] > heights[j] {
				temp = names[i]
				names[i] = names[j]
				names[j] = temp

				temp2 = heights[i]
				heights[i] = heights[j]
				heights[j] = temp2
			}
		}
	}

	// reverse the order of the names
	for i := 0; i < len(names)/2; i++ {
		temp = names[i]
		names[i] = names[len(names)-i-1]
		names[len(names)-i-1] = temp
	}

	return names
}
