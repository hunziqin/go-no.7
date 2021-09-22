package main

import "fmt"

func main() {
	var a [15]int
	for i, _ := range a {
		a[i] = i
		fmt.Printf("%d\n", a[i])
	}
}
