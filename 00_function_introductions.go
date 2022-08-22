package main

import (
  "fmt"
)
 
// Function:
func add(value_a int, value_b int) int {
    return value_a + value_b
}

// Call the function within the main() function:
func main() {
  result:= add(2, 3)
  fmt.Println("The result of the function 'add':", result)
}
