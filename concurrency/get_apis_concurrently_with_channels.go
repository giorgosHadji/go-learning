package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()
	// When you create a channel by using the make() function, you create an unbuffered channel, which is the default behavior.
	// Unbuffered channels block the sending operation until there's someone ready to receive the data.
	ch := make(chan string)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api, ch)
	}
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	// Here we can see it finishes in <Done! It took 0 seconds!> super fast, as we didnt actually waited for the results of the goroutines.
}
func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}
