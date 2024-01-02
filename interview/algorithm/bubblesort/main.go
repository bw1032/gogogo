package main

import (
	"fmt"
)

func main() {
	var s = []int{8, 2, 3, 1, 7, 6, 4, 5, 9}
	//bubbleSort1(s) //[1 2 3 4 5 6 7 8 9]
	bubbleSort2(s) //[9 8 7 6 5 4 3 2 1]
}

func bubbleSort1(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] { //找后面小的放在前面
				arr[i], arr[j] = arr[j], arr[i]
			}
			fmt.Println(arr)
		}
		fmt.Println("---------------------------")
	}
}

func bubbleSort2(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			fmt.Println(arr)
		}
		fmt.Println("---------------------------")
	}
}
