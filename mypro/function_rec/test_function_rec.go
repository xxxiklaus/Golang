/*
	golang函数递归

函数内部调用自身的函数称为递归函数
使用递归函数最重要的三点：
1.递归就是自己调用自己
2.必须先定义函数的退出条件，没有退出条件，递归将成为死循环
3.golang递归函数很可能产生一大堆goroutine，也可能出现栈空间内存溢出问题
*/
package main

import "fmt"

//阶乘

func f1() {
	var s int = 1
	// 5! = 5x4x3x2x1
	for i := 1; i <= 5; i++ {
		s *= i //*= 先相乘后赋值
	}
	fmt.Printf("s: %v\n", s)
}

func f2(a int) int {
	if a == 1 { // 1.退出条件
		return 1
	} else {
		return a * f2(a-1) //自己调用自己
	}
}

//斐波那契数列 f(n)=f(n-1)+f(n-2)且f(2)=f(1)=1
// 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987

func f3(n int) int {
	//退出点判断
	if n == 1 || n == 2 { // ||逻辑运算符OR 这里即n=1或者n=2 即退出 返回1
		return 1
	}
	//递归表达式
	return f3(n-1) + f3(n-2)
}

func main() {
	// f1()
	// i := f2(8)
	// fmt.Printf("i: %v\n", i)
	r := f3(10)
	fmt.Printf("r: %v\n", r)

}
