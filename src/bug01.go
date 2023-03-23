package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.

// Issue: Channels are meant to send and receive values from Goroutines.
// Since there are no Goroutines the channels wait indefinitely
// for input/output.
// Solution: Use a Goroutine with an anonymous function
// to send the value to the channel.
func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}
