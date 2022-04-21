package main

import "fmt"

func main() {
	messages := make(chan string)
	for i := 0; i < 5; i++ {
		go broker(messages)
		fmt.Println(<-messages)
	}
}

func broker(messages chan string) {
	messages <- "ping"
}
