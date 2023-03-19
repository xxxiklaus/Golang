// golang切片的遍历
// 切片的遍历和数组的遍历非常类似，可以使用for循环索引遍历，或者for range循环

package main

import "fmt"

func test1() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	l := len(s1)
	for i := 0; i < 1; i++ {
		fmt.Printf("s1[%v]: %v\n", i, s1[i])
		fmt.Printf("l: %v\n", l)

	}
}

func test2() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	for i, v := range s1 {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)

	}

}

func main() {
	// test1()
	test2()

}
