package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
// Issue: Channels are meant to send and receive values from Goroutines. 
// But, the main Goroutine is not a Goroutine. So, it cannot send or receive values from a channel.
// Instead, the channels wait indefinitely for input/output.
// Solution: Use a Goroutine to send the value to the channel. 
func main() {
	ch := make(chan string)
	go func ()  {
		ch <- "Hello world!"
	}()

	fmt.Println(<-ch)
	close(ch)
}
