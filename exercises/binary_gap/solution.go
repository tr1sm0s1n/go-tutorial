package main

import "fmt"

func Solution(N int) int {
	binary := fmt.Sprintf("%b", N) // Convert N to its binary representation
	maxGap := 0                    // To track the longest gap
	currentGap := 0                // To track the current gap length
	inGap := false                 // To track if we're inside a gap

	for _, ch := range binary {
		if ch == '1' {
			if inGap {
				// End of a gap
				if currentGap > maxGap {
					maxGap = currentGap
				}
				currentGap = 0 // Reset the current gap
			}
			inGap = true // Start tracking a gap
		} else if inGap {
			// Inside a gap, count the zeros
			currentGap++
		}
	}

	return maxGap
}

func main() {
	fmt.Println(Solution(1041)) // Output: 5
	fmt.Println(Solution(32))   // Output: 0
	fmt.Println(Solution(625))  // Output: 3
}
