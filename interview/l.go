package main

func L() {
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
}
