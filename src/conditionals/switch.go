package main

import (
	"fmt"
)

func main() {

	switch "docker" {
	case "linux":
		fmt.Println("\nHere are some linux courses")
	case "docker":
		fmt.Println("\nHere are some docker courses")
	case "windows":
		fmt.Println("\nHere are some window courses")
	default:
		fmt.Println("\nSorry, we couldn't find a match!")
	}
}
