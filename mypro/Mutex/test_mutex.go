/*
	golang并发编程之Mutex互斥锁实现同步

除了使用channel实现同步之外,还可以使用Mutex互斥锁的方式实现同步
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var m int = 100

var lock sync.Mutex

var wt sync.WaitGroup

func add() {
	defer wt.Done()
	lock.Lock() //加锁
	m += 1
	fmt.Printf("m++: %v\n", m)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock() //执行完 解锁
}

func sub() {
	defer wt.Done()
	lock.Lock() //加锁
	time.Sleep(time.Millisecond * 2)
	m -= 1
	fmt.Printf("m--: %v\n", m)
	lock.Unlock() // 执行完 解锁
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		wt.Add(1)
		go sub()
		wt.Add(1)
	}
	wt.Wait()
	fmt.Printf("end m: %v\n", m)

}
