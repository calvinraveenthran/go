package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/Users/calvinraveenthran/Documents/code/go/src/introducing-go/hello-world.go")

	if err != nil {
		fmt.Println("Error returned was:", err)
		return
	}

	fmt.Println("No error was returned")
}
