package main

import "fmt"

var a, b, c int

func main() {
	fmt.Scan(&a, &b, &c)
	if a < b && b < c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
