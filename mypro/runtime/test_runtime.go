/*
	golang并发编程之runtime包

runtime包里面定义了一些协程管理相关的api
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)

	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
		time.Sleep(time.Millisecond * 100)
	}
}

/* func show() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit()  // 退出当前协程
		}

	}
} */

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1) //修改为1 查看效果 定义使用CPU核心数
	go a()
	go b()
	time.Sleep(time.Second)
	/* 	go show()
	   	time.Sleep(time.Second) */
}

/* func show(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}
func main() {
	go show("java") //子协程来运行
	for i := 0; i < 5; i++ {
		// runtime.Gosched() 让出CPU时间片，重新等待安排任务
		runtime.Gosched() //我有权利执行任务了，让给你（其他子协程来执行）
		fmt.Printf("\"golang\": %v\n", "golang")
	}
	fmt.Println("end...")

}
*/
