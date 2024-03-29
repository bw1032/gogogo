package main

import "fmt"

/*
对照表:
	字节数	Unicode			十进制			UTF-8 编码
	1		000000-00007F	0-127			0xxxxxxx
	2		000080-0007FF	128-2047		110xxxxx 10xxxxxx
	3		000800-00FFFF	2048-65535		1110xxxx 10xxxxxx 10xxxxxx
	4		010000-10FFFF	65536-1114111	11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
*/

/*
南 unicode编码:

	二进制	0101001101010111
	十六进制  5357 (\u5357)
	十进制	21335

在Unicode编码中，"\u"表示一个Unicode字符的转义序列的开始。而"5357"是该Unicode字符的十六进制编码值。
具体来说，"\u5357"表示Unicode字符U+5357，即汉字"南"。
*/
func main() {

	d := "南"
	d1 := '南'
	fmt.Println([]byte(d)) //[229 141 151]
	//二进制：0101001101010111 -> utf8: 11100101 10001101 10010111 -> 十进制：229 141 151
	fmt.Println([]rune(d)) //[21335]
	//0101001101010111 -> 十进制：21335
	fmt.Println(byte(d1)) //87
	/*
		字符常量使用单引号'表示，而不是双引号"。当使用单引号包围一个字符时，它实际上表示一个整数值，即该字符的Unicode码点。
		对于字符常量s := '南'，它会被解释为Unicode码点U+5357（即汉字"南"的Unicode编码）。
		因此，执行fmt.Println(byte(s))会将该Unicode码点转换为byte类型，并打印输出对应的字节值。
		然而，在Go语言中，byte类型实际上是uint8的别名，表示一个8位无符号整数。
		对于Unicode码点U+5357，它超出了byte类型的范围（0~255），因此会发生截断。
		所以，输出结果87表示将U+5357转换成byte类型时发生了截断，最低有效字节为0x57，即十进制的87。
		这是一个错误的结果，因为它只是取了截断后的低8位，而没有正确地表示整个Unicode字符。
		正确的方式是使用字符串来表示字符常量，如s := "南"，或者使用rune类型来表示字符，如s := rune('南')。
		这样可以保留完整的Unicode字符，并正确处理其编码。
	*/
	fmt.Println(rune(d1)) //21335
	//二进制：0101001101010111 -> 十进制：21335

	return
}
