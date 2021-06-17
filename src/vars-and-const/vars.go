package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Nigel"
	course := "Docker Deep Dive"
	module := "3.2"

	fmt.Println("Name is set to:", name, "and is of type ", reflect.TypeOf(name))
	fmt.Println("Module is set to:", module, "and is of type ", reflect.TypeOf(module))
	fmt.Println("Course is set to:", course, "and is of type ", reflect.TypeOf(course))
}
