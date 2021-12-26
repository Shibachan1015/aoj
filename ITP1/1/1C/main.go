package main 

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Printf("%v %v\n", a*b, 2*(a+b))
}
