package main

import "fmt"

func main() {
	slice1 := make([]int, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * 5
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
