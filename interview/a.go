package main

import "fmt"

//普通类型(int,string等) 和 数组 是值传递，拷贝一份
//map、slice、channel、指针类型这些特殊类型 是引用传递

func A() {

	c := []int32{1, 2}
	d := c
	d[0] = 3
	fmt.Println(c)     //[3 2]   切片c赋值给d，修改d[0]的值，c[0]也跟着变了，说明切片是引用传值

	a := 1
	b := a
	b = b + 1
	fmt.Println(a) //1  给b赋值 a的值没有变化 说明是值传递

    e := [3]int32{1,2}
    f := e
    f[0] = 3
    fmt.Println(e)

	fmt.Println(fmt.Sprintf("类型:%T,地址:%p", c, &c)) //类型:[]int32,地址:0xc000118078
	fmt.Println(fmt.Sprintf("类型:%T,地址:%p", a, &a)) //类型:int,地址:0xc00011a160

}
