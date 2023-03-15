package main

import (
	"fmt"
	"time"
	"sync"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
// Issue: The main function finishes before the Print function, so the Print function never gets to print 11.
// Solution: Make the program wait for the Print function to finish before exiting.
func main() {
	ch := make(chan int)
	wgp := new(sync.WaitGroup)
	wgp.Add(1)
	
	go Print(ch, wgp)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wgp.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wgp *sync.WaitGroup) {

	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	wgp.Done()
}
