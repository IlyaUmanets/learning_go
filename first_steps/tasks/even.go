package main

import "fmt"

/*
Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.
For example half(1) should return (0, false) and half(2) should return (1, true).
*/

func main() {
	fmt.Println(half(10))
}

func half(number int) bool {
  half_number := number / 2
  result := false
  if (half_number % 2 == 0) {
    result = true
  }
  return result
}
