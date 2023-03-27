package main

import (
	"fmt"
)

func showline(sequence []int) {
	for _, n := range sequence {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()

}

func main() {
	n := 9

	for i := 0; i < n; i++ {
		fibonacci_number_at_i := fibonacci_reqursion(i)
		fmt.Printf("%d, ", fibonacci_number_at_i)
	}
	fmt.Println()
	numbs := fibonacci(n)
	fmt.Println(numbs)
	totalsum := sum(numbs)
	fmt.Println("the total from", n, "numbers is", totalsum)

}

func fibonacci(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			sequence[i] = i
		} else {
			sequence[i] = sequence[i-1] + sequence[i-2]
			showline(sequence)
		}
	}
	return sequence
}

func fibonacci_reqursion(n int) int {
	var value int

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	prev := fibonacci_reqursion(n - 1)
	prev_prev := fibonacci_reqursion(n - 2)

	value = prev + prev_prev

	return value
}

func sum(numbs []int) int {
	total := 0
	for _, n := range numbs {
		total += n
		// fmt.Println("sum func space", total)
	}
	return total
}
