package main

import (
	"fmt"
)

//runtime hchan结构:
//type hchan struct {
//	qcount   uint           // 数组长度，即已有元素个数
//	dataqsiz uint           // 数组容量，即可容纳元素个数
//	buf      unsafe.Pointer // 数组地址
//	elemsize uint16         // 元素大小
//	closed   uint32
//	elemtype *_type // 元素类型
//	sendx    uint   // 下一次写下标位置
//	recvx    uint   // 下一次读下标位置
//	recvq    waitq  // 读等待队列
//	sendq    waitq  // 写等待队列
//	lock     mutex
//}

func G() {
	c1 := make(chan int)
	var c2 chan int
	fmt.Println(fmt.Sprintf("%T,%+v,%p", c1, c1, c1)) //chan int,0xc0001022a0,0xc0001022a0
	fmt.Println(fmt.Sprintf("%T,%+v,%p", c2, c2, c2)) //chan int,<nil>,0x0

	//读未初始化的chan：
	<-c2 //fatal error: all goroutines are asleep - deadlock! //goroutine 1 [chan receive (nil chan)]:

	//写未初始化的chan:
	c2 <- 1 //fatal error: all goroutines are asleep - deadlock! //goroutine 1 [chan send (nil chan)]:


	close(c1)

	//读已关闭的chan
	num, ok := <-c1
	fmt.Println(num, ok) //0 false
	//如果 chan 关闭前，buffer 内有元素还未读 , 会正确读到 chan 内的值，且返回的第二个 bool 值（是否读成功）为 true
	//如果 chan 关闭前，buffer 内有元素已经被读完，chan 内无值，接下来所有接收的值都会非阻塞直接成功，
	//返回 channel 元素的零值， 但是第二个 bool 值一直为 false


	//写已关闭的chan
	c1 <- 1 //panic: send on closed channel //goroutine 1 [running]:
}
