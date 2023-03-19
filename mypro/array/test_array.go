// golang数组
// 数组是相同类型的一组数据的集合，数组一旦定义长度不能修改，
// 数组可以通过下标（或者叫索引）来访问元素
// golang数组的定义
// 语法 var variable_name [SIZE] variable_type
// variable_name 数组名称
// SIZE 数组长度，必须是常量
// variable_type 数组保存的元素类型

package main

import "fmt"

func test1() {
	var a1 [2]int              //定义一个int类型的数组a1，长度是2
	var a2 [3]string           //定义一个string类型的数组a2，长度是3
	var a3 [2]bool             //布尔类型
	fmt.Printf("a1: %T\n", a1) //%T打印类型
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2) // 打印a1，a2本身默认值 整型则均为0，字符串则为空的字符串
	fmt.Printf("a3: %v\n", a3)
}

func test2() {
	//数组的初始化
	//初始化列表 即将值写在大括号里面
	var a1 = [3]int{1, 2}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("a3: %v\n", a3)
}

//数组长度可以省略，使用...代替，更加初始化值得数量自动推断

func test3() {
	var a1 = [...]int{1, 2, 3}
	var a2 = [...]string{"hello", "go", "hello", "world"}
	var a3 = [...]bool{false, true, true, false, false}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}

//指定索引值的方式来初始化
// 可以通过指定所有的方式来初始化，未指定所有的默认未零值
func test4() {
	var a1 = [...]int{0: 1, 2: 2}
	var a2 = [...]string{1: "hello", 2: "go"}
	var a3 = [...]bool{2: true, 5: false}
	a4 := [...]int{1, 2}

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()

}
