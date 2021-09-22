package main

import "fmt"

func f(a [3]int)  { fmt.Printf("%p\n", &a) }
func fp(a [3]int) { fmt.Printf("%p", &a) }

func main() {
	var ar [3]int
	f(ar)
	fp(ar)
}
