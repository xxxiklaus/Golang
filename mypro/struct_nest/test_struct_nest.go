/*
	golang嵌套结构体

go语言没有面向对象编程思想，也没有继承关系
但是可以通过结构体嵌套来实现这种效果
*/
package main

import "fmt"

func main() {

	type Dog struct {
		name  string
		age   int
		color string
	}

	type Person struct {
		dog  Dog
		name string
		age  int
	}

	dog := Dog{
		name:  "大黄",
		age:   2,
		color: "yellow",
	}

	per := Person{
		dog:  dog,
		name: "klaus",
		age:  22,
	}

	fmt.Printf("per.dog.name: %v\n", per.dog.name) //使用（.） 运算符直接访问嵌套里想要知道的信息

	fmt.Printf("per: %v\n", per)

}
