package main

import "fmt"

func main() {
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)

	// adding elements
	studentsAge1 := make(map[string]int)
	studentsAge1["john"] = 32
	studentsAge1["bob"] = 31
	fmt.Println(studentsAge1)

	// Notice that we've added items to a map that was initialized. But if you try to do the same with a nil map, you'll get an error.
	// panic: assignment to entry in nil map
	var studentsAgeError map[string]int
	studentsAgeError["john"] = 32
	studentsAgeError["bob"] = 31
	fmt.Println(studentsAgeError)
	// To avoid problems when you add items into a map, make sure you create an empty map (not a nil map) by using the make function like we showed in the previous code snippet.
	// This rule only applies when you add items. If you run lookups, deletes, or loop operations in a nil map, Go won't panic.
}
