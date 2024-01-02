package main

import (
	"fmt"
	"runtime"
	"time"
)

func I() {

	goroute(1)
	// 单核 输出: 先出输出了等号 再输出了协程里的打印
	// 因为
	//那么go出去的goroutine到底什么时候能执行到呢？golang的调度器是非抢占式的，在GPM的架构里，
	//Waiting态的goroutine（在LRQ或者GRQ中等待执行）
	//必须要等Executing态的goroutine主动退出执行，才能绑定到M上执行。
	//主动退出的方式有很多，比如time.Sleep()，比如runtime.Gosched()。

	// ==================
	// 30
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// .....

	//goroute(2)
	// 双核 输出: 协程无需等主进程退出执行  就可开始执行  所以等号 和 协程打印  不固定输出

	// 3
	// 4
	// ==================
	// 40
	// 5
	// 6
	// 7
	// ...

}

func goroute(cpu int) {

	runtime.GOMAXPROCS(cpu)

	for i := 0; i <= 10; i++ {
		go func(a int) {
			fmt.Println(a)
		}(i)
	}

	fmt.Println("==================")
	time.Sleep(3 * time.Second)

}
