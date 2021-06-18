package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "calvin raveenthran"

	fmt.Println(converter(module, author))
}

func converter(module string, author string) (s1 string, s2 string) {

	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
