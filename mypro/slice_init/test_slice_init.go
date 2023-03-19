package main

import "fmt"

//切片

func test1() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("s1: %v\n", s1) // 直接初始化
	s2 := s1[0:3]              // 前边是闭区间 后边是开区间
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4)
	s5 := s1[:] //使用数组初始化
	fmt.Printf("s5: %v\n", s5)
}

//数组
/* 切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片，切片表达式中
的low和high表示一个索引范围（即左边包含右边不包含 左闭右开），得到的切片长度等于
high-low，容量等于得到的切片的底层数组容量
*/
func test2() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6}
	a2 := a1[:3]
	fmt.Printf("a2: %v\n", a2)
	a3 := a1[3:]
	fmt.Printf("a3: %v\n", a3)
	a4 := a1[:]
	fmt.Printf("a4: %v\n", a4)

}

func main() {
	// test1()
	test2()

}
