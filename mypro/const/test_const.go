// go语言常量
/*常量，就是在程序编译阶段就确定下来的值，而程序在运行时则无法改变该值，在Go程序中
常量可以是数值类型(包括整型、浮点型和复数型)、布尔类型、字符串类型等。*/

//定义常量的语法
/*定义一个常量使用const关键字，语法格式：
const constanName [type]= value
const: 定义常量关键字
constanName： 常量名称
type：常量类型
value：常量的值 */
package main

import "fmt"

func main() {

	/* 	const (
	   		a = 100
	   		b = 200
	   	)
	   	fmt.Printf("a: %v\n", a)
	   	fmt.Printf("b: %v\n", b)

	   	const w, h = 200, 300
	   	fmt.Printf("w: %v\n", w)
	   	fmt.Printf("h: %v\n", h) */
	//iota
	/*iota比较特殊，可以被认为是一个可被编译器修改的常量，它默认开始值是0，每调用一次加1，
	  直到遇到const关键字时被重置为0.*/
	/*   const (
			a1 = iota  //0
			a2 = iota //i++
			a3 = iota
		)
		fmt.Printf("a1: %v\n", a1)
		fmt.Printf("a2: %v\n", a2)
		fmt.Printf("a3: %v\n", a3)
	 const PI float32 = 3.14 //PI=3.15
	 const PI2 = 3.1415
	 fmt.Printf("PI: %v\n", PI) */
	//   使用_跳过某些值
	/* 	const (
	   		a1 = iota // 0
	   		_        // 1
	   		a2 = iota // 2
	   	)
	   	fmt.Printf("a1: %v\n", a1)
	   	fmt.Printf("a2: %v\n", a2) */
	//iota 声明中间插队
	const (
		a1 = iota
		a2 = 100
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

}
