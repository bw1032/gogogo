package main

import (
	"testing"
)

// go test -run TestC
func TestA(t *testing.T) {
	//方法或函数调用传参数, 数组用值传递, map、slice、channel、指针类型这些特殊类型是引用传递。
	A()
}

func TestB(t *testing.T) {
    //字符串相关
	B()
}

func TestC(t *testing.T) {
	//翻转中文字符串
	C()
}

func TestD(t *testing.T) {
	//nil切片和空切片
	D()
}


func TestE(t *testing.T) {
	//reflect 反射
	E()
}


func TestG(t *testing.T) {
	//select case
	G()
}

func TestH(t *testing.T) {
	//runtime
	H()
}

func TestI(t *testing.T) {
	//gorutinue
	I()
}

func TestJ(t *testing.T) {
	//gorutinue
	J()
}

func TestK(t *testing.T) {
	//gorutinue
	K()
}