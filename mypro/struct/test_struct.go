/*
	golang结构体

go语言没有面向对象的概念了，但是可以用结构体来实现
面向对象编程的一些特性，eg：继承、组合等特性
golang结构体定义 语法结构如下
type struct_variable_type struct{  // type:结构体定义关键字

		member definition;  //struct_variable_type: 结构体类型名称
		member definition;  //struct: 结构体定义关键字
		...
		member definition; // member definition: 成员定义
	}
*/
package main

import "fmt"

//自定义了一个类型
type Person struct { //定义一个人的结构体
	id    int
	name  string
	age   int
	email string
}

type Customer struct {
	id, age     int
	name, email string
}

func main() {
	//匿名结构体 如果结构体是临时使用，可以不用起名字，直接使用 eg：
	var klaus struct {
		id, age     int
		name, email string
	}
	klaus.id = 818
	klaus.age = 22
	klaus.name = "klaus"
	klaus.email = "pilw818@gmail.com"
	fmt.Printf("klaus: %v\n", klaus)
	//访问结构体成员 可以使用点运算符(.)来访问结构体成员
	/* 	var klaus Person
	   	klaus.id = 101
	   	klaus.name = "klaus"
	   	klaus.age = 22
	   	klaus.email = "pilw818@gmail.com"
	   	fmt.Printf("klaus: %v\n", klaus) */

	/* 	var klaus Person
	   	fmt.Printf("klaus: %v\n", klaus)
	   	var sui Customer
	   	fmt.Printf("sui: %v\n", sui) */

}
