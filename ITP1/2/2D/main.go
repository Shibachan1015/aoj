package main

import (
	"fmt"
)

var (
	w int
	h int
	x int
	y int
	r int
)

func main() {
	fmt.Scanf("%d %d %d %d %d", &w, &h, &x, &y, &r)
	if x - r >= 0 && x + r <= w && y - r >= 0 && y + r <= h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
