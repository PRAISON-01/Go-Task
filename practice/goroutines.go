package main

import (
	"fmt"
	"time"
)

// pinger runs an infinite loop sending "ping" to the channel
func pinger(channel chan string) {
	for i := 0; ; i++ {
		channel <- "ping" // Blocks here until printer reads it
	}
}

// printer runs an infinite loop reading from the channel and printing every second
func printer(channel chan string) {
	for {
		msg := <-channel // Blocks here until pinger sends data
		fmt.Println(msg)
		time.Sleep(time.Second * 1) // Dictates the 1-second pace of the application
	}
}

func ponger(channel chan string) {
	for i := 0; ; i++ {
		channel <- "pong"
	}
}

func main() {
	// Create the unbuffered string communication channel
	var c chan string = make(chan string)

	// Launch both concurrent workers in the background
	go ponger(c)
	go pinger(c)
	go printer(c)

	// Keep the main goroutine alive until you hit 'Enter' in your terminal
	var input string
	fmt.Scanln(&input)
}
