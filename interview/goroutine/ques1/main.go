package main

import (
	"fmt"
)

// 两个协程交替打印123456
func main() {
	sign1 := make(chan int)
	sign2 := make(chan int)
	num := make(chan int)
	end := make(chan struct{})
	endnum := 6

	go func() {
		for {
			a := <-sign2
			num <- a
			if a == endnum {
				end <- struct{}{}
				return
			}
			sign1 <- a + 1
		}
	}()

	go func() {
		for {
			a := <-sign1
			num <- a
			if a == endnum {
				end <- struct{}{}
				return
			}
			sign2 <- a + 1
		}
	}()

	go func() {
		for {
			fmt.Print(<-num)
		}
	}()

	sign2 <- 1
	<-end
	close(end)
}
