package main

//参考
//https://www.cnblogs.com/bigdataZJ/p/go-channel-deadlock.html

func A() {
	//fatal error: all goroutines are asleep - deadlock!

	//死锁1：一个通道在一个主go程里,同时读和写
	//ch := make(chan string)
	//ch <- "1"
	//fmt.Println(<-ch)


	//死锁2：go程开启之前使用通道
	//ch := make(chan string)
	//ch <- "1" 此处死锁 优于go程之前使用通道
	//go func() {
	//	fmt.Println(<-ch)
	//}()
	//ch <- "1" 此处不死锁
	//time.Sleep(2 * time.Second)


	//死锁3：在通道1中调用通道2 在通道2中调用通道1

	//死锁4：超过channel缓存继续写入导致死锁
	//ch:= make(chan string, 2)
	//ch <- "1"
	//ch <- "2"
	//ch <- "3"
	//fmt.Println(<-ch)


	//死锁5：向已关闭的channel写入 不会导致死锁 但会导致panic
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(" err : ", r)  //捕获打印panic
	//	}
	//}()
	//
	//ch := make(chan string)
	//close(ch)
	//ch <- "1"


}
