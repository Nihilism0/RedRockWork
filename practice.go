package main

import (
	"fmt"
)

func main() {
	var a string
	ret := 1
	count := 0
	fmt.Scanf("%s", &a)
	b := []rune(a)
	for i := 0; i < len(b)/2; i++ {
		if b[i] != b[len(b)-i-1] {
			ret = 0
		}
	}
	if ret == 0 {
		fmt.Println("false")
	} else {
		for _, c := range b {
			count++
			fmt.Printf("%c", c)
			if count == len(b)/2 {
				break
			}
		}
	}
}
