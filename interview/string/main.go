package main

import "fmt"

func main() {
	s := "你abc好!"

	var src = []rune(s)
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}

	fmt.Println(s)
	fmt.Println(string(src))
}
