package simbol

import "fmt"

type point struct {
	x, y int
}

func Simbol() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
}
