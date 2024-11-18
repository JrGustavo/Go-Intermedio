package main

import "testing"

func TestSum(t *testing.T) {

	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 1},
		{1, 2, 2},
		{2, 2, 2},
		{5, 2, 5},
	}

	for _, table := range tables {
		max := GetMax(table.x, table.y)
		if max != table.n {
			t.Errorf("Max of (%d,%d) was incorrect, got: %d, want: %d.", table.x, table.y, max, table.n)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		x int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, table := range tables {
		fib := Fibonacci(table.x)
		if fib != table.n {
			t.Errorf("Fibonacci of (%d) was incorrect, got: %d, want: %d.", table.x, fib, table.n)
		}
	}
}
