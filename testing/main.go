package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}