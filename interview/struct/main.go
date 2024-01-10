package main

import (
	"fmt"
	"unsafe"
)

/*
		struct能不能比较？ 很显然这句话包含了两种情况：
		   同一个struct的两个实例能不能比较？
		   两个不同的struct的实例能不能比较？

		在分析上面两个问题前，先跟大家梳理一下golang中，哪些数据类型是可比较的，哪些是不可比较的：
		   可比较：Integer，Floating-point，String，Boolean，Complex(复数型)，Pointer，Channel，Interface，Array
		   不可比较：Slice，Map，Function

		同一个struct的2个实例能不能比较 == !=
		   答案：可以能、也可以不能

	   两个不同的struct的实例能不能比较 == !=
	       答案：可以能、也可以不能

		   如果结构体的所有成员变量都是可比较的，那么结构体就可比较
		   如果结构体中存在不可比较的成员变量，那么结构体就不能比较
		   结构体之间进行转换需要他们具备完全相同的成员(字段名、字段类型、字段个数)

*/

/*
	GO语言中int类型的大小是不确定的，与具体的平台有关系一般来说，int在32位系统中是4字节，在64位系统中是8字节使用自动推导类型初始化一个整数，默认为int类型。

	int8: -128 ~ 127
	int16: -32768 ~ 32767
	int32: -2147483648 ~ 2147483647
	int64: -9223372036854775808 ~ 9223372036854775807
	uint8: 0 ~ 255
	uint16: 0 ~ 65535
	uint32: 0 ~ 4294967295
	uint64: 0 ~ 18446744073709551615
*/

func main() {

	//内存对齐规则： 1 成员对齐规则  2 整体对齐规则

	var a int
	var a1 int8
	var a2 int32
	var a3 int64
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(a1), unsafe.Sizeof(a2), unsafe.Sizeof(a3))
	//8 1 4 8

	type A struct {
		A int8
	}
	fmt.Println(unsafe.Sizeof(A{})) //1

	type B struct {
		B  int8  //1
		B1 int64 //8
		B2 int32 //4
	}
	b := B{}
	fmt.Println(unsafe.Sizeof(b))                                                 //24
	fmt.Println(unsafe.Alignof(b.B), unsafe.Sizeof(b.B), unsafe.Offsetof(b.B))    //1 1 0
	fmt.Println(unsafe.Alignof(b.B1), unsafe.Sizeof(b.B1), unsafe.Offsetof(b.B1)) //8 8 8
	fmt.Println(unsafe.Alignof(b.B2), unsafe.Sizeof(b.B2), unsafe.Offsetof(b.B2)) //4 4 16

	//0 1 2 3 4 5 6 7  8  9 10 11 12 13 14 15 16  17 18 19 20 21 22 23
	//B                B1                     B2         |
	type C struct {
		C  int8
		C1 int32
		C2 int64
	}
	c := C{}
	fmt.Println(unsafe.Sizeof(c))                                                 //16
	fmt.Println(unsafe.Alignof(c.C), unsafe.Sizeof(c.C), unsafe.Offsetof(c.C))    //1 1 0
	fmt.Println(unsafe.Alignof(c.C1), unsafe.Sizeof(c.C1), unsafe.Offsetof(c.C1)) //4 4 4
	fmt.Println(unsafe.Alignof(c.C2), unsafe.Sizeof(c.C2), unsafe.Offsetof(c.C2)) //8 8 8

	//0 1 2 3 4 5 6 7  8  9 10 11 12 13 14 15
	//B       B1       B2
}
