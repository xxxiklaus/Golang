/*
	golang并发编程之协程

Golang中的并发是函数相互独立运行的能力.Goroutines是并发运行的函数
Golang提供了Goroutines作为并发处理操作的一种方式
创建一个携程非常简单，就是在一个任务函数前面添加一个go关键字:go task()
*/
package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 8; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("java")   // go 启动了一个协程来执行
	go showMsg("golang") //2
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("main end...") //3 主函数退出 程序就结束了  即主死随从

}
