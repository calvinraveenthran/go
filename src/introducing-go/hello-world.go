package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello from", runtime.GOOS) // This tells you what operating system you are using!
}
