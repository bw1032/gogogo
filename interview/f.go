package main

import "fmt"

func F() {

	//这个代码会造成死循环吗？
	s := []int{1, 2}
	for _, v := range s {
		s = append(s, v)
		fmt.Printf("len(s)=%v\n", len(s))
	}
	//len(s)=3
	//len(s)=4

	//不会死循环，for range其实是golang的语法糖，在循环开始前会获取切片的长度 len(切片)，然后再执行len(切片)次数的循环。
}
