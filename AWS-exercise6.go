package main

import "fmt"

func main() {
	var ar = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	/*var a = ar[1:1]
	fmt.Printf("len(a) is %d\n", len(a))*/
	var a = ar[1:2]
	fmt.Printf("len(a) is %d", len(a))
}
