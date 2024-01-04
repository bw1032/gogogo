package main

import (
	"fmt"
)

func main() {
	sign1 := make(chan int)
	sign2 := make(chan int)
	go func() {
		close(sign1)
	}()

	go func() {
		close(sign2)
	}()

	select {
	case a := <-sign1:
		fmt.Println(a)
	case a := <-sign2:
		fmt.Println(a)
	}
}
