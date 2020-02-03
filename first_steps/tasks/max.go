package main

import "fmt"

/*
Write a function with one variadic parameter that finds the greatest number in a list of numbers.
*/

func main() {
	fmt.Println(max(1, 8, 3))
}

func max(args ...int) int {
  result := 0
  for _, el := range args {
    if el > result {
      result = el
    }
  }
  return result
}
