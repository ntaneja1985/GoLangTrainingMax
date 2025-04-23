package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	//We can pass this channel to different go routines
	fmt.Println("Starting...")
	go slowGreet("Hello slow greet", done)
	go Greet("Hello fast", done)

	for range done {
		//fmt.Println(doneChan)
	}
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	//This arrow points in the direction of where data should flow
	doneChan <- true
	close(doneChan)
}

func Greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}
