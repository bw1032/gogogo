package main

import (
	"fmt"
	"time"
)

/*
根据官方文档
A "select" statement chooses which of a set of possible send or receive operations will proceed.
It looks similar to a "switch" statement but with the cases all referring to
communication operations.
其实很类似于switch语句，只不过case语句上有些差别——case只能引用通信操作。
什么是“通信操作”呢？就可以理解为对chan的读取或写入操作啦！


其表达意义在于，当多个需要从多个chan中读取或写入时，会先轮询一遍所有的case，
然后在所有处于就绪（可读/可写）的chan中随机挑选一个进行读取或写入操作，并执行其语句块。
如果所有case都未就绪，则执行default语句，如未提供default语句，则当前协程被阻塞。

*/

func K() {

	fmt.Println("====")
	a := <- time.After(3 * time.Second)
	fmt.Println("====1")
	fmt.Println(a)

	chan1 := make(chan string)
	chan2 := make(chan string)

	select {
	case v, k := <-chan1:
		fmt.Println(k)
		fmt.Println(v)
	case v, k := <-chan2:
		fmt.Println(k)
		fmt.Println(v)
	//default:
	//	fmt.Println("end")
	}

}
