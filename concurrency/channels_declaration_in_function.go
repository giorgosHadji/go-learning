package main

import "fmt"

func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
	// you get < invalid operation: ch <- "Bye!" (send to receive-only type <-chan string)>, as you declared the channel as a receive only channel.
	// ch <- "Bye!"
}

func main() {
	ch := make(chan string, 1)
	send(ch, "Hello World!")
	read(ch)
}
