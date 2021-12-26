package main

import "fmt"

func main() {
	var a, b int16
	fmt.Scan(&a, &b)

	if a < b {
		fmt.Println("a < b")
	} else if  a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a > b")
	}
}
