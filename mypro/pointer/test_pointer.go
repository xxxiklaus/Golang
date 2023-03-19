/*
	golang指针

Go语言中的函数传参都是值拷贝，当我们遇到想要修改某个变量的时候，我们可以创建一个指向该
变量地址的指针变量。传递数据使用指针，而无需拷贝数据
类型指针不能进行偏移和运算
Go语言中的指针操作非常简单，只需要牢记两个符号 &（取地址） 和*（根据地址取值）
指针地址和指针类型
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置.Go语言中使用&字符在变量
前面对变量进行取地址操作。Go语言中的值类型(int、float、bool、string、array、struct)
都有对应的指针类型，eg:"*int、*int64、*string"
指针语法
一个指针变量指向了一个值的内存地址。（也就是我们声明了一个指针之后，可以像变量赋值一样
把一个值的内存地址放到指针当中）
类似于变量和常量，在使用指针前你需要声明指针，指针语法如下：
var var_name *var-type
var-type : 指针类型
var_name : 指针变量名
* : 用于指定变量作为一个指针
*/
package main

import "fmt"

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip) // nil
	fmt.Printf("ip: %T\n", ip)

	var i int = 100
	ip = &i // 取地址
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip) // *ip取值

	var sp *string                //声明指针变量
	var s string = "hello golang" //声明实际变量
	sp = &s
	fmt.Printf("sp: %T\n", sp)
	fmt.Printf("sp: %v\n", *sp)

	var bp *bool
	var b bool = true
	fmt.Printf("bp: %T\n", bp)
	bp = &b
	fmt.Printf("bp: %v\n", *bp)

}
