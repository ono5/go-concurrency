package main

import "fmt"

func getMsg(ch1 chan<- string) {
	defer close(ch1)
	ch1 <- "Hello"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	defer close(ch2)
	message := <-ch1
	ch2 <- message
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go getMsg(ch1)
	go relayMsg(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}
}
