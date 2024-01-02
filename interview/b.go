package main

import (
	"fmt"
	"unicode/utf8"
	"golang.org/x/exp/utf8string"
)

func B() {

	//string的索引操作返回的是byte（或uint8），
	//如想获取字符可使用for range，也可使用unicode/utf8包和golang.org/x/exp/utf8string包的At()方法
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	var c = "a你好"
	fmt.Println(c[0], c[1], string(c[1])) //97 //228 //ä
	fmt.Println(utf8string.NewString(c).At(0), utf8string.NewString(c).At(1)) //97 20320
	fmt.Println(string(utf8string.NewString(c).At(1))) //你
	fmt.Println(utf8string.NewString(c).Slice(1,2))  //你

	//utf8.RuneCountInString(a) 字符长度  len(a) 字节长度
	var a = "a"
	fmt.Println(utf8.RuneCountInString(a), len(a)) //1 //1
	var b = "你好"
	fmt.Println(utf8.RuneCountInString(b), len(b)) //2 //6
}
