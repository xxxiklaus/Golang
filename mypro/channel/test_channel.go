/*
	Golang并发编程之通道channel

Go提供了一种称之为通道的机制,用于在goroutines直接共享数据.当您作为goroutines执行并发活动时
需要在goroutines之间共享资源或数据,通道充当goroutine之间的管道并提供一种机制来保证同步交换
需要在声明通道时指定数据类型.我们可以共享内置、命名、结构和引用类型的值和指针.数据在通道上传递：
在任何给定时间只有一个goroutine可以访问数据项,因此设计不会发生数据竞争.
根据数据交换的行为,有两种类型的通道:无缓冲通道和缓冲通道.无缓冲通道用于执行goroutine之间的同步
通信,而缓冲通道用于执行异步通信.无缓冲通道保证在发送和接收发生的瞬间执行两个goroutine之间的交换.
缓冲通道没有这样的保证
通道由make函数创建,该函数指定chan关键字和通道元素类型
语法：
Unbuffered := make(chan int) //整型无缓冲通道
buffered := make(chan int,10) // 整型有缓冲通道

使用内置函数make创建无缓冲和缓冲通道.make的第一个参数需要关键字 cahn,然后是通道允许交换的数据类型

将值发送到通道的代码块 需要使用 <- 运算符 语法：
goroutine1 := make(chan string,5) //字符缓冲通道
goroutine1 <- "Australia" //通过通道发送字符串

从通道接收值的代码块 语法:
data := <-goroutine1 // 从通道接收数字字符串

无缓冲通道
在无缓冲通道中,在接收到任何值之前都没有能力保存它,在这钟类型的通道中,发送和接收goroutine在任何
发送或接收操作完成之前的同一时刻都准备就绪.如果两个goroutine没有在同一时刻准备好,则通道会让执行
其各自发送或接收操作的goroutine首先等待.同步是通过通道上发送和接收之间交互的基础,没有另一个就不可能发生

缓冲通道
在缓冲通道中，有能力在接收到一个或多个值之前保存它们.在这种类型的通道中,不要强制goroutine在同一时刻
准备好执行发送和接收.当发送或接收阻塞时也有不同的条件.只有当通道中没有要接收的值时,接收才会阻塞.仅当
没有可用缓冲区来放置正在发送的数值时,发送才会阻塞

通道的发送和接收特性

1.对于同一个通道,发送操作之间是互斥的,接收操作之间也是互斥的
2.发送操作和接收操作中对元素的值处理都是不可分割的
3.发送操作在完全完成之前会被阻塞,接收操作也是如此
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建int类型通道，只能传入int类型值
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}

func main() {
	//从通道接收值
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")

}
