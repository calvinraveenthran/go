package main

import (
	"fmt"
)

func main() {

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than the second course")
	} else if secondRank < firstRank {
		fmt.Println("\nSecond course is doing better than the first course")
	} else {
		fmt.Println("\nThey have the same rank")
	}
}
