package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//for_append_()
	//for_append_2()
	//nil_empty_()
	//pointer()
}

/**/
func for_append_() {
	//代码会造成死循环吗？
	s := []int{1, 2}
	for _, v := range s {
		s = append(s, v)
		fmt.Printf("len(s)=%v\n", len(s))
	}
	//不会，for range是go的语法糖，循环开始前会获取切片的长度 len(切片)，然后再执行len(切片)次数的循环。
}

/**/
func for_append_2() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("value :%d value-Addr:%x ElemAddr：%x\n", value, &value, &slice[index])
	}
	//第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份 副本。
	//range创建的是每个元素的副本，而不是返回对该元素的引用
}

/**/
func nil_empty_() {

	var s1 []int         //nil切片
	s2 := make([]int, 0) //空切片
	s3 := make([]int, 0) //空切片
	s4 := make([]int, 1) //初始值[]int{0}
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

/**/
func pointer() {

	X := func(arr []int) {
		arr[0] = 2
	}

	Y := func(arr []int) {
		arr = append(arr, 2)
		arr[0] = 2
	}

	var a = []int{1}
	X(a)
	fmt.Println(a) //[2] 函数内和函数外 是同一个切片底层的数组指针

	var b = []int{1}
	Y(b)
	fmt.Println(b) //[1] 函数内切片已扩容  不是同一个数组
}

/**/
func expansion() {
	//元素小于1024
	//expansion_1()
	//expansion_2()
	//元素大于等于1024
	//expansion_3()
	//
	var s = []string{"1"}
	fmt.Println(unsafe.Pointer(&s[0]), s, len(s), cap(s), float64(cap(s))*1.25)
	expansion_4(s)
	fmt.Println(unsafe.Pointer(&s[0]), s, len(s), cap(s), float64(cap(s))*1.25)
}

func expansion_1() {
	var s []string
	for i := 1; i <= 9; i++ {
		s = append(s, "1")
		fmt.Println(unsafe.Pointer(&s), unsafe.Pointer(&s[0]), len(s), cap(s))
	}
	fmt.Println(s)

	//0xc000092060 0xc000088020 1 1
	//0xc000092060 0xc0000983c0 2 2 //容量扩展两倍
	//0xc000092060 0xc0000cc000 3 4 //容量扩展两倍
	//0xc000092060 0xc0000cc000 4 4
	//0xc000092060 0xc0000ce000 5 8 //容量扩展两倍
	//0xc000092060 0xc0000ce000 6 8
	//0xc000092060 0xc0000ce000 7 8
	//0xc000092060 0xc0000ce000 8 8
	//0xc000092060 0xc0000d0000 9 16 //容量扩展两倍
	//[1 1 1 1 1 1 1 1 1]

}

func expansion_2() {
	s := make([]string, 2, 2)
	for i := 1; i <= 9; i++ {
		s = append(s, "1")
		fmt.Println(unsafe.Pointer(&s), unsafe.Pointer(&s[0]), len(s), cap(s))
	}
	fmt.Println(s)

	//0xc000008078 0xc00004c040 3 4
	//0xc000008078 0xc00004c040 4 4
	//0xc000008078 0xc000078000 5 8 //容量扩展两倍
	//0xc000008078 0xc000078000 6 8
	//0xc000008078 0xc000078000 7 8
	//0xc000008078 0xc000078000 8 8
	//0xc000008078 0xc00007a000 9 16 //容量扩展两倍
	//0xc000008078 0xc00007a000 10 16
	//0xc000008078 0xc00007a000 11 16
	//[  1 1 1 1 1 1 1 1 1]

}

func expansion_3() {
	var s []string
	for i := 1; i <= 1025; i++ {
		s = append(s, "1")
		if len(s) == cap(s) {
			fmt.Println(len(s), cap(s), float64(cap(s))*1.25)
		}
	}

	//1 1 1.25
	//2 2 2.5
	//4 4 5
	//8 8 10
	//16 16 20
	//32 32 40
	//64 64 80
	//128 128 160
	//256 256 320
	//512 512 640
	//848 848 1060

}

func expansion_4(s []string) {

	for i := 1; i <= 10; i++ {
		s = append(s, "1")
		if len(s) == cap(s) {
			fmt.Println(unsafe.Pointer(&s[0]), len(s), cap(s), float64(cap(s))*1.25)
		}
	}
}
