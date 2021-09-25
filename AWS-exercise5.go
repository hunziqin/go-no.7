package main

import "fmt"

func main() {
	var a int = 10
	f(a)
}

func f(n int) {
	var fibonacci []int = make([]int, 10)
	for i := 0; i < len(fibonacci); i++ {
		if i <= 1 {
			fibonacci[i] = 1
		} else {
			fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
		}
	}
	for i := 0; i < len(fibonacci); i++ {
		fmt.Printf("fibonacci at %d is %d\n", i, fibonacci[i])
	}
	return
}
