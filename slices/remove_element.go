package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove_index := 2

	if remove_index < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove_index])

		letters = append(letters[:remove_index], letters[remove_index+1:]...)

		fmt.Println("After", letters)
	}

}
