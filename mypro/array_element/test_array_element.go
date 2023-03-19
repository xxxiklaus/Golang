// golang访问数组元素
// 可以通过下标的方式来访问数组元素。数组最大下标为数组长度-1
// 大于这个下标会发生数组越界
package main

import (
	"fmt"
)

func test1() {
	var a [2]int
	a[0] = 100
	a[1] = 200

	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Printf("a[1]: %v\n", a[1])

	//修改 a[0] a[1]
	a[0] = 1
	a[1] = 2
	fmt.Println("------------")

	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Printf("a[1]: %v\n", a[1])

	// a1[3] = 1000  // 数组长度越界 会直接报错

}

func test2() {
	//数组的长度
	var a1 = [3]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))

	var a2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len(a2): %v\n", len(a2))

}

// 根据数组长度遍历数组
// 可以根据数组长度，通过for循环的方式来遍历数组，数组的长度可以使用len函数获得
func test3() {
	//1.根据长度和下标
	var a1 = [3]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
}

//2.for range

func test4() {
	var a1 = [3]int{1, 2, 3}
	for i, v := range a1 {
		fmt.Printf("a1[%v]: %v\n", i, v)

	}
}

func main() {
	// test1()
	// test2()
	// test3()
	test4()

}
