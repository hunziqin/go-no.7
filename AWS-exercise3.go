package main

import "fmt"

func main() {
	var a [50]int
	for i, _ := range a {
		if i <= 1 {
			a[i] = 1
		} else {
			a[i] = a[i-1] + a[i-2]
		}
		fmt.Printf("fibinacci(%d) is %d\n", i, a[i])
	}
}
