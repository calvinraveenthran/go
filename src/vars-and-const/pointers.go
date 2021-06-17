package main

import (
	"fmt"
)

func main() {

	name := "Calvin"
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course)

	course = changeCourseByValue(course)

	fmt.Println("\nHi", name, "you're now watching", course)

	changeCourseByReference(&course)

	fmt.Println("\nHi", name, "Finally you're now watching", course)
}

func changeCourseByValue(course string) string {

	course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", course)

	return course
}

func changeCourseByReference(coursePtr *string) {

	*coursePtr = "First Look: Native Docker Clustering 2"

	fmt.Println("\nTrying to change your course to", *coursePtr)

}
