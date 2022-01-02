package main

import "fmt"

var a, b, c, x int

func main () {
	fmt.Scan(&a, &b, &c)
	for i := a; i <= b; i++ {
		if c % i == 0 {
			x++
		}
	}
	fmt.Println(x)
}
