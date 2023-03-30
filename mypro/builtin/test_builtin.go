/*
	golang标准库builtin

builtin包提供了一些类型声明,变量,常量声明,还有一些遍历函数
这个包不需要导入,这些变量和函数就可以直接使用
*/
package main

import "fmt"

func testAppend() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)
	s1 := []int{4, 5, 6}
	i2 := append(s, s1...)
	fmt.Printf("i2: %v\n", i2)
}

func testLen() {
	s := "Hello Golang"
	fmt.Printf("len(s): %v\n", len(s))
	s1 := []int{1, 2, 3}
	fmt.Printf("len(s1): %v\n", len(s1))
}

func testPrint() {
	name := "klaus"
	age := 22
	print(name, " ", age, "\n")
	fmt.Println("------")
	println(name, " ", age)

}

func testPanic() {
	defer fmt.Println("panic 后 我还会执行...")
	panic("坏 寄了...")
	fmt.Println("end...")
}

/* new和make的区别 */
// 1.make只能用来分配及初始化类型为slice,map,chan的数据;new可以分配任意类型的数据
// 2.new分配返回的是指针,即类型 *T; make返回引用,即 T;
// 3.new分配的空间被清零,make分配后,会进行初始化

func testNew() {
	b := new(bool)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %v\n", *b)
	i := new(int)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("i: %v\n", *i)
	s := new(string)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("s: %v\n", *s)

}

//内建函数make(T,args)与new(T)的用途不同.它只是用来创建slice,map,channel,并且返回
//一个初始化的(而不是置零),类型为T的值(而不是*T).之所以有所不同,是因为这三个类型的背后
//引用了使用前必须初始化的数据结构.例如,slice是一个三元描述符,包含一个指向数据(在数组中)
//的指针,长度,以及容量,在这些项被初始化之前,slice都是nil的.对于slice,map,和channel.
//make初始化这些内部数据结构,并准备号可用的值

/* eg: make([] int, 10, 100)
分配一个有100个int的数组,然后创建一个长度为10,容量为100的slice结构,该slice引用包含前10个元素的
数组.对应的,new([]int)返回一个指向新分配的,被置零的slice结构体的指针,即指向值为nil的slice指针 */

func testMakeVsNew() {
	var p *[]int = new([]int) //allocates slice structure; *p == nil; rarely useful
	fmt.Printf("p: %v\n", p)
	// var v []int = make([]int, 100) //the slice v now refers to a new array of 100 ints
	v := make([]int, 10) // Idiomatic:习惯的做法
	fmt.Printf("v: %v\n", v)

}

func main() {
	// testAppend()
	// testLen()
	// testPrint()
	// testPanic()
	// testNew()
	testMakeVsNew()

}
