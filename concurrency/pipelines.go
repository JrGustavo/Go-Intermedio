package main

import "fmt"

func Generator(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in chan int, out chan int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(c chan int) {
	for value := range c {
		fmt.Print(value)
	}
}

func main() {
	generator := make(chan int)
	double := make(chan int)
}
