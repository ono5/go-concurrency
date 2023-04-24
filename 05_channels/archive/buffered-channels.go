package main

import "fmt"

func main() {
	ch := make(chan int, 6)
	go func() {
		defer close(ch)

		// send all iterator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %d\n", v)
	}
}

/*
ch := make(chan int)
Sending: 0
Sending: 1
Received: 0
Received: 1
Sending: 2
Sending: 3
Received: 2
Received: 3
Sending: 4
Sending: 5
Received: 4
Received: 5
*/

/*
ch := make(chan int, 6)
Sending: 0
Sending: 1
Sending: 2
Sending: 3
Sending: 4
Sending: 5
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
*/
