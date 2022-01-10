package main

import "fmt"

func main () {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		x := i
		if i % 3 == 0 {
			fmt.Printf(" %d ", i)
			continue
		}
		for x > 0 {
			if x % 10 == 3 {
				fmt.Printf(" %d", i)
				break
			}
			x /= 10
		}
	}
	fmt.Println("")
}
