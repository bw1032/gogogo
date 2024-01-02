package main

import (
	"fmt"
	"runtime"
	"sync"
)

func H() {
	run1()
	//B: 9
	//A: 10
	//A: 10
	//A: 10
	//A: 10
	//A: 10
	//A: 10
	//A: 10
	//A: 10
	//A: 10
	//A: 10
	//B: 0
	//B: 1
	//B: 2
	//B: 3
	//B: 4
	//B: 5
	//B: 6
	//B: 7
	//B: 8

	//GOMAXPROCS = 1; goroutine顺序执行 先打印B: 9 跟golang的调度有关 调度的源码里有个next_goroutine的指针，
	//会取最后一次开启的goroutine的地址，所以住routine被等待后的next 就是你最后一次开启的那个goroutine

	run2()
	//BB: 9
	//AA: 10
	//AA: 10
	//AA: 10
	//AA: 10
	//AA: 10
	//BB: 4
	//BB: 0
	//BB: 1
	//BB: 2
	//BB: 3
	//BB: 6
	//BB: 5
	//AA: 10
	//AA: 10
	//AA: 10
	//AA: 10
	//BB: 7
	//BB: 8
	//AA: 10

	//GOMAXPROCS = 4; goroutine无序执行
}

func run1() {
	runtime.GOMAXPROCS(1)
	var wait sync.WaitGroup
	wait.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A:", i)
			wait.Done()
		}()
	}
	for j := 0; j < 10; j++ {
		go func(j int) {
			fmt.Println("B:", j)
			wait.Done()
		}(j)
	}
	wait.Wait()
}

func run2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wait sync.WaitGroup
	wait.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("AA:", i)
			wait.Done()
		}()
	}
	for j := 0; j < 10; j++ {
		go func(j int) {
			fmt.Println("BB:", j)
			wait.Done()
		}(j)
	}
	wait.Wait()
}
