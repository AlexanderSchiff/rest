package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	array := os.Args[1:]
	fmt.Print("Sum = ")
	fmt.Print(AddArray(array))
	fmt.Println()
}

// AddArray adds up ints in a string array
func AddArray(array []string) int {
	sum := 0
	for _, str := range array {
		val, err := strconv.Atoi(str)
		if err == nil {
			sum += val
		} 
	}
	return sum
}