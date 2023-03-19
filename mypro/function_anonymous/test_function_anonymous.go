/*
	goalng匿名函数

go语言函数不能嵌套，但是在函数内部可以定义匿名函数，实现以下简单功能调用
所谓匿名函数就是没有名称的函数，语法格式
func (参数列表)(返回值) （当然可以既没有参数也没有返回值）
*/
package main

import "fmt"

func test1() {
	name := "klaus "
	age := "20 "
	email := "pilw818@gmail.com"

	//在函数内部做一些运算

	f1 := func() string {
		return name + age + email
	}

	msg := f1()
	fmt.Printf("msg: %v\n", msg)
}

func main() {
	//匿名函数
	/* 	max := func(a, b int) int {
	   		if a > b {
	   			return a
	   		} else {
	   			return b
	   		}
	   	}
	   	f := max(1557, 4396)
	   	fmt.Printf("f: %v\n", f) */
	//自己执行
	/* func(a, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Printf("max: %v\n", max)
	}(1557, 4396) */
	test1()
}
