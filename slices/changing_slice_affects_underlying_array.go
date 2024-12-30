package main

import "fmt"

//when you change an element from a slice, you're changing the underlying array too. Any other slices that refer to the same underlying array will be affected.
// thats why you needs "copies"
func main() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 := letters[1:4]

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
}
