package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func D() {

	var s1 []int  //nil切片
	s2 := make([]int,0) //空切片
	s3 := make([]int,0) //空切片
	s4 := make([]int,1) //初始值[]int{0}
	fmt.Printf("s1 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("s2 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	fmt.Printf("s3 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
	fmt.Printf("s4 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	//s1 pointer:{Data:0 Len:0 Cap:0}
	//s2 pointer:{Data:824633974392 Len:0 Cap:0}
	//s3 pointer:{Data:824633974392 Len:0 Cap:0}
	//s4 pointer:{Data:824633974392 Len:1 Cap:1}

	//结论 nil切片没有地址  空切片指向同一个地址

	//s1[1] = 1 //panic: runtime error: index out of range [1] with length 0
	s1 = append(s1, 1) // √
	fmt.Printf("s1 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	s1 = append(s1, 1) // √
	fmt.Printf("s1 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	//s1 pointer:{Data:824633762352 Len:1 Cap:1}
	//s1 pointer:{Data:824633762384 Len:2 Cap:2}


	//s2[1] = 1 //panic: runtime error: index out of range [1] with length 0
	s2 = append(s2, 1) // √
	fmt.Printf("s2 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	s2 = append(s2, 1) // √
	fmt.Printf("s2 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	//s2 pointer:{Data:824633762424 Len:1 Cap:1}
	//s2 pointer:{Data:824633762464 Len:2 Cap:2}

}
