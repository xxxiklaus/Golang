/*
	golang闭包

闭包可以理解成定义在一个函数内部的函数。在本质上，闭包是将内部函数
和外部函数连接起来的桥梁.或者说是函数和其引用环境的组合体
闭包指的是一个函数与其相关联的引用环境组合而成的实体。简单来说，闭包=函数+引用环境(牢记)
首先我们来看一个例子
*/
package main

import (
	"fmt"
	"strings"
)

//返回一个函数

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y // 相加后再赋值
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func cal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i //相加后再赋值  base是环境变量
		return base
	}

	sub := func(i int) int {
		base -= i //相减后再赋值
		return base
	}
	return add, sub
}

func main() {
	/* 	f := add()
	   	r := f(10)
	   	fmt.Printf("r: %v\n", r)
	   	r = f(20)
	   	fmt.Printf("r: %v\n", r)
	   	r = f(30)
	   	fmt.Printf("r: %v\n", r)
	   	fmt.Println("-------")
	   	f1 := add()
	   	r = f1(100)
	   	fmt.Printf("r: %v\n", r)
	   	r = f1(200)
	   	fmt.Printf("r: %v\n", r) */
	/* jpgFunc := makeSuffixFunc(".jpg\n")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Printf(jpgFunc("test"))
	fmt.Printf(txtFunc("test")) */
	f1, f2 := cal(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))

}
