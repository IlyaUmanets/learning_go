package main

import "fmt"

/*
sum is a function which takes a slice of numbers and adds them together.
What would its function signature look like in Go?
*/

func main() {
	fmt.Println(sum(1, 2, 3))
}

func sum(numbers ...int) int {
	total := 0
  for _, number := range numbers {
    total += number
  }
	return total
}
