package main

import "fmt"

func main() {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Printf("len(a) is %d and cap(a) is %d\n", len(a), cap(a))
	fmt.Println(a, " ")
	a = a[0:4]
	fmt.Printf("len(a) is %d and cap(a) is %d\n", len(a), cap(a))
}
