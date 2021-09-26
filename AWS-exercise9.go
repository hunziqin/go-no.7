package main

import "fmt"

func main() {
	a := "goland"
	index := 2
	s1, s2 := f(a, index)
	fmt.Printf("The string %s split at position %d is: %s / %s\n", a, index, s1, s2)
}

func f(str string, i int) (s1, s2 string) {
	b := []byte(str)
	s1 = string(b[:i])
	s2 = string(b[i:])
	return
}
