package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a string
	ret := 1
	count := 0
	fmt.Scanf("%s", &a)
	lens := len(a)
	b := []rune(a)
	for _, c := range a {
		if unicode.Is(unicode.Han, c) {
			lens -= 2
		}
	}
	for i := 0; i < lens/2; i++ {
		if b[i] != b[lens-i-1] {
			ret = 0
		}
	}
	if ret == 0 {
		fmt.Println("false")
	} else {
		for _, c := range b {
			count++
			fmt.Printf("%c", c)
			if count == lens/2 {
				break
			}
		}
	}

}
