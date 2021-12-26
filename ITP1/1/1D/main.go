package main

import "fmt"

var overall int 
var h, m, s int

func main() {
	fmt.Scan(&overall)
	s = overall%60
	m = (overall/60)%60
	h = (overall/60)/60
	fmt.Printf("%v:%v:%v\n", h, m, s)
}
