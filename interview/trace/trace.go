package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

/*
适用场景： 找出程序在一段时间内正在做什么
不适用场景： 运行缓慢的函数，或者分析大部分CPU时间花费在哪里，CPU有专门的工具 go tool pprof

go tool trace tracefile

	View trace：查看跟踪
	Goroutine analysis：Goroutine 分析
	Network blocking profile：网络阻塞概况
	Synchronization blocking profile：同步阻塞概况
	Syscall blocking profile：系统调用阻塞概况
	Scheduler latency profile：调度延迟概况
	User defined tasks：用户自定义任务
	User defined regions：用户自定义区域
	Minimum mutator utilization：最低 Mutator 利用率
*/
func main() {

	runtime.GOMAXPROCS(1) //设置可以同时执行的cpu的最大数量为1个
	f, _ := os.Create("trace")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()

	//自定义一个任务
	ctx, task := trace.NewTask(context.Background(), "customerTask")
	defer task.End()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ { //启动10个协程，模拟做任务
		go func(num string) {
			defer wg.Done()
			trace.WithRegion(ctx, num, func() {
				var sum, j int64
				// 模拟执行任务
				for ; j < 500000000; j++ {
					sum += j
				}
				fmt.Println(num, sum)
			})
		}(fmt.Sprintf("num_%02d", i))
	}
	wg.Wait()
}
