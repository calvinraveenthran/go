package main

import (
	"fmt"
	"reflect"
)

var (
	name   string
	course string
	module float64
)

func main() {
	fmt.Println("Name is set to:", name, "and is of type ", reflect.TypeOf(name))
	fmt.Println("Modile is set to:", module, "and is of type ", reflect.TypeOf(module))
}
