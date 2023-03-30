/* golang并发编程之channel遍历 */
package main

import "fmt"

var c = make(chan int)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) //调用close关掉程序
	}()

	for {
		v, ok := <-c
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			break
		}
	}

	/* 	for v := range c {
		fmt.Printf("v: %v\n", v)
		for range 最简洁

	} */

	/* 	for i := 0; i < 10; i++ {
		r := <-c
		fmt.Printf("r: %v\n", r)
		用for循环的方式读取r

	} */

}
