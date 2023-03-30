/*
	golang并发编程之原子操作详解

atomic提供的原子操作能够确保任一时刻只有一个goroutine对变量进行操作
善用atomic能够避免程序中出现大量的锁操作
atomic常见操作有:
增减、载入 load 、比较并交换 cas、 交换、 存储 write;
*/
package main

import (
	"fmt"
	"sync/atomic"
)

func test_add_sub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -1)
	fmt.Printf("i: %v\n", i)

	var j int64 = 200
	atomic.AddInt64(&j, 1)
	fmt.Printf("j: %v\n", j)
}

func test_load_store() {
	// Load 载入操作能够保证原子的读变量的值,当读取的时候,任何其他CPU操作都无法对该变量进行读写,其实现机制受到底层硬件的支持
	// Store 存储 此类操作确保了写变量的原子性,避免其他操作读到了修改变量过程中的脏数据
	var i int32 = 100
	atomic.LoadInt32(&i) //read
	fmt.Printf("i: %v\n", i)

	atomic.StoreInt32(&i, 200) //write
	fmt.Printf("i: %v\n", i)
}

func test_cas() {
	//cas (CompareAndSwap) 该操作在进行交换前首先确保变量的值未被更改，即仍然保持参数old所记录的值,满足此前提下才
	// 进行交换操作.CAS的做法类似于操作数据库时常见的乐观锁机制
	var i int32 = 100

	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("i: %v\n", i)
}

// Swap 相对于CAS,明显此类操作更为暴力直接,并不管变量的旧值是否改变,直接赋予新值然后返回被替换的值
func main() {

}
