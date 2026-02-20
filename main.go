package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan any) {
	fmt.Println("Hello!", phrase)
	doneChan <- "channel"
}

func slowGreet(phrase string, doneChan chan any) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)

}

func main() {
	// go greet("Nice to meet you!")
	done := make(chan any)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!")
	<-done
	<-done

	for range done {
		fmt.Println(<-done)
	}
}
