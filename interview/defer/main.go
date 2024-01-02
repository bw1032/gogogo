package main

import "fmt"

/*
	    https://www.cnblogs.com/l199616j/p/15500631.html#_label0
		defer、return、返回值三者的执行是：
			return最先执行，先将结果写入返回值中（即赋值）；
			接着defer开始执行一些收尾工作；
			最后函数携带当前返回值退出（即返回值）。
*/

type P struct {
	Age int
}

func main() {
	//case1
	p := &P{10}
	defer fmt.Print(p.Age)
	defer func(p *P) {
		fmt.Print(p.Age)
	}(p)
	defer func() {
		fmt.Print(p.Age)
	}()
	p.Age = 20
	//280 280 28

	//case2
	//fmt.Println(A())
}

func A() int {
	a := 0
	defer func() {
		a += 1
	}()
	return a
}
