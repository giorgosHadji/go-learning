package main

import "fmt"

// The capacity of a slice only tells you how much you can extend a slice.

//For instance, if you try to print something like fmt.Println(quarter2[3]), you'll get the following error: panic: runtime error: index out of range [3] with length 3.
func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter2 := months[3:6]
	quarter2Extended := quarter2[2:6]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
}
