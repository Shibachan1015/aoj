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
				if i == 0 || j == 0 || i == H-1 || j == W-1 {
					var sh string = "#"
					fmt.Printf("%s", sh)
				} else {
					var dot string = "."
					fmt.Printf("%s", dot)
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
