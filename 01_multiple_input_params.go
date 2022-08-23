package main

import (
	"fmt"
)

func add_values(init_value int, next_values ...int) []int {
	result := make([]int, 0, len(next_values))
	for _, value := range next_values {
		result = append(result, init_value+value)
    // Adds up the first value with the following value(s).
	}
	return result
}

func main() {
	fmt.Println(add_values(1))
	fmt.Println(add_values(1, 2))
	fmt.Println(add_values(3, 4, 9, 7, 1))
  // Works also with different data structures
  set_values := []int{28, 90, 13, 29}
  fmt.Println(add_values(90, set_values...))
}
