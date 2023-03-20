/*
	golang接口

接口像是公司里面的领导，他会定义一些通用规范，只设计规范，而不实现规范
golang的接口，是一种新的类型定义，它把所有的具有共性的方法定义在一起
任何其他类型只要实现了这些方法就是实现了这个接口
*/
/*  接口语法格式
//定义接口
type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
    ...
	method_namen [return_type]

}

// 定义结构体
type struct_name struct {
	//  variables
}

//实现接口方法
func (struct_name_variable struct_name)method_name1()[return_type] {
	// 方法实现
}

func (struct_name_variable struct_name)method_namen()[return_type] {
	// 方法实现
}
// 在接口定义中定义，若干个空方法，这些方法都具有通用性   */

package main

import (
	"fmt"
)

//定义一个USB接口，有读read和写write两个方法，再定义一个电脑Computer和手机Mobile来实现这个接口

// Writer Reader 接口名称一般以er结尾
type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

type Mobile struct {
	model string
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read...")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write...")
}

func (m Mobile) read() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("read...")
}

func (m Mobile) write() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("write...")
}

func main() {
	c := Computer{
		name: "Macbook air",
	}

	c.read()
	c.write()

	m := Mobile{
		model: "Iphone 14 pro",
	}

	m.read()
	m.write()

}

// 注：实现接口必须实现接口中的所有方法，否之会报
/* 一个变量如果实现了接口中规定的所有的方法,那么这个变量就实现了这个接口
可以称为这个接口的变量
*/
