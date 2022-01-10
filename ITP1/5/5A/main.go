package main

import (
	"fmt"
)

func main() {
	var H int = 0
	var W int = 0	
	for {
		fmt.Scanf("%d %d", &H, &W)
		if H == 0 && W == 0 {
			break
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
			
			var sh string = "#"
			fmt.Printf("%s", sh)
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
