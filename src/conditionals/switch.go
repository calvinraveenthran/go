package main

import (
	"fmt"
)

func main() {

	switch "docker" {
	case "docker":
		fmt.Println("\nHere are some docker courses")
		fallthrough //prints the guy below it.
	case "linux":
		fmt.Println("\nHere are some linux courses")
		fallthrough
	case "windows":
		fmt.Println("\nHere are some window courses")
		fallthrough
	default:
		fmt.Println("\nSorry, we couldn't find a match!")
	}
}
