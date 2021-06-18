package main

import (
	"fmt"
)

func main() {

	firstRank := 2
	secondRank := 2

	if firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than the second course")
	} else if secondRank < firstRank {
		fmt.Println("\nSecond course is doing better than the first course")
	} else {
		fmt.Println("\nThey have the same rank")
	}
}
