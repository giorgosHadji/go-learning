package main

import "fmt"

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

// Go convention dictates that if any method of a struct has a pointer receiver, all methods of that struct must have a pointer receiver.
//  Even if a method of the struct doesn't need it. --> Althought it doesnt create an error or a warning if you don't do it, so its not actually enforced.
func (t *triangle) doubleSize() {
	t.size *= 2
}
func main() {
	t := triangle{3}
	t.doubleSize()
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
}
