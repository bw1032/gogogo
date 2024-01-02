package main

import (
	"fmt"
)

func J() {
	c := make(chan string)

	go func() {
		c <- "hello"
		close(c)
	}()

	for v := range c { //for range可以自动监测到channel关闭，然后自动退出
		fmt.Println(v)
	}

}
