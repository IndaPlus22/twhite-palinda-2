// Stefan Nilsson 2013-03-13

/*
What happens if you switch the order of the statements wgp.Wait() 
and close(ch) in the end of the main function?
The function will end before the Waitgroup is finished.
The Goroutines won't have any channels to send to so the program ends.

What happens if you move the close(ch) from the main function and 
instead close the channel in the end of the function Produce?
The first Produce routine will close the channel and the other channels 
won't have any channels to send to. As a result, an error will occur

What happens if you remove the statement close(ch) completely?
The program will still work as the channels shut down at the end
of the main function. Closing is only necessary if you want the 
receiver to know that no more data will be coming, like in a 
a range or for loop.

What happens if you increase the number of consumers from 2 to 4?
More threads will be created, thus the program will be faster

Can you be sure that all strings are printed before the program stops?
No. The program will stop when all the producers are done, but since 
there is no Waitgroup for the consumers it's possible there are still
strings in the channel when the program terminates

*/

// This is a testbed to help you understand channels better.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgp)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 < x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
