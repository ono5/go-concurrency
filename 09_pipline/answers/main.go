package main

import "fmt"

// generator - converts a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	// ch := generator(2, 3)
	// out := square(ch)

	// for n := range out {
	// 	fmt.Println(n)
	// }
	for n := range square(generator(2, 3)) {
		fmt.Println(n)
	}
}
