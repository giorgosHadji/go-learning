package main

import (
	"fmt"

	"github.com/gh/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	// .\geometry_main.go:12:24: t.size undefined (cannot refer to unexported field size)
	// size as its not capitalized is considered a private member, which you cannot access from a different package.
	// fmt.Println("Size", t.size)
	fmt.Println("Perimeter", t.Perimeter())
}
