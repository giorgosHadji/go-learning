package main

import (
	"fmt"
)

func main() {
	// total := calculator.Sum(3, 5)
	// fmt.Println(total)
	// fmt.Println("Version: ", calculator.Version)
	// fmt.Println(quote.Hello())
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		switch {
		case i%3 == 0:
			{
				fmt.Println("fuzz")
			}
		}
	}
}
