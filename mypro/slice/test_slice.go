//golang切片
// 数组是固定的长度，可以容纳相同数据类型的元素的集合。当长度固定时
// 使用还是带来一些限制，eg：我们申请的长度太大浪费内存，太小又不够用
// 鉴于上述原因我们需要使用golang的切片，可以把切片理解为，可变长度的数组
// 其实它底层就是使用数组实现的，增加了自动扩容功能。切片（Slice）是一个拥有相同
// 类型元素的可变长度的序列
// golang切片的语法
/* 声明一个切片和声明一个数组类似，只要不添加长度即可
var identifier []type
切片是引用类型，可以使用make函数来创建切片:
var slice1 []type = make([]type, len)
简写: slice := make([]type, len)
也可以指定容量，其中capacity为可选参数.
make([]T, length, capacity) // 这里len是数组的长度并且也是切片的初始长度 */
package main

import "fmt"

//直接声明，还未初始化
func test1() {
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

//make
func test2() {
	var s1 = make([]int, 2)
	fmt.Printf("s1: %v\n", s1)
}

//切片拥有自己的长度和容量，可以通过使用内置的len（）函数求长度
// 使用内置的cap（）函数求切片的容量

func test3() {
	var s1 = []int{1, 2, 3}
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))

	fmt.Printf("s1[0]: %v\n", s1[0])
}

func main() {
	// test1()
	// test2()
	test3()

}
