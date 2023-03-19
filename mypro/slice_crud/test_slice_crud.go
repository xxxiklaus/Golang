// go语言切片元素的添加和删除copy
// 切片是一个动态数组，可以使用append（）函数添加元素，go语言中并没有删除
// 切片元素的专用方法，我们可以使用切片本身的特性来删除元素。由于切片是引用
// 类型，通过赋值的方式，会修改原有内容，go提供了copy（）函数来拷贝切片

// 增删改查&copy

package main

import "fmt"

//add
func test1() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}

//delete

func test2() {
	var s1 = []int{1, 2, 3, 4, 5}
	//i 2
	s1 = append(s1[:2], s1[3:]...) // 先取到索引0、1的数 即{1，2}  再取索引2之后的数 即{4，5}
	// a = append(a[:index],a[index+1:]...)  从而达到删除索引2对应的数{3}

	fmt.Printf("s1: %v\n", s1)
}

//update

func test3() {
	var s1 = []int{1, 2, 3, 4, 5}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}

// query

func test4() {
	var s1 = []int{1, 2, 3, 4, 5}
	var key = 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}

	}
}

//copy

func test5() {
	var s1 = []int{1, 2, 3, 4, 5}
	// var s2 = s1
	var s2 = make([]int, 5)
	copy(s2, s1)
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}
