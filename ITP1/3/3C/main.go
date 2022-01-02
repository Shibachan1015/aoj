package main

import "fmt"

var (
	x, y, tmp int
)

func main () {
	for i := 1; ;i++ {
		fmt.Scan(&x, &y)
		if x == 0 && y == 0 {
			break
		}
		if x > y {
			tmp = y
			y = x
			x = tmp
		} 
		fmt.Printf("%d %d\n", x, y)
	}
}
