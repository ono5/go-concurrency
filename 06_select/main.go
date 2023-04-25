package main

import (
	"fmt"
	"time"
)

func main() {
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch1 <- "one"
	// }()

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch2 <- "two"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case c1 := <-ch1:
	// 		fmt.Println(c1)
	// 	case c2 := <-ch2:
	// 		fmt.Println(c2)
	// 	}
	// }
	// select {
	// case m1 := <-ch1:
	// 	fmt.Println(m1)
	// case m2 := <-ch2:
	// 	fmt.Println(m2)
	// case <-time.After(1 * time.Millisecond):
	// 	fmt.Println("TIME OUT")
	// }

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no message received")
		}

		fmt.Println("processing...")
		time.Sleep(1500 * time.Millisecond)
	}
}
