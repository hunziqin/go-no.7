package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'd'}
	fmt.Printf("%s \n", b[1:4])
	fmt.Printf("%s \n", b[:2])
	fmt.Printf("%s \n", b[2:])
	fmt.Printf("%s ", b[:])
}
