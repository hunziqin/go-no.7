package main

const (
	WIDTH  = 1920
	HEIGHI = 1080
)

type pixel int

var screen [WIDTH][HEIGHI]pixel

func main() {
	for y := 0; y < HEIGHI; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
}
