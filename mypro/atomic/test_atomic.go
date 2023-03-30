/* golang并发编程之原子变量的引入 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// var i =100
var i int32 = 100

// 没有加锁 引起线程并发之间出现的问题
var lock sync.Mutex

func add() {
	atomic.AddInt32(&i, 1) //cas compare and swap  比较和交换 old new
	/* 	lock.Lock()
	   	i++
	   	lock.Unlock() */
}

func sub() {
	atomic.AddInt32(&i, -1)
	/* 	lock.Lock()
	   	i--
	   	lock.Unlock() */
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second * 2)

	fmt.Printf("i: %v\n", i)

}
