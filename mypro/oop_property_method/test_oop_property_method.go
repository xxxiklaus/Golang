/*
	golang模拟OOP的属性和方法

golang没有面向对象的概念,也没有封装的概念,但是可以通过结构体struct
和函数绑定来实现OOP的属性和方法等特性,接收receiver方法
*/
package main

import "fmt"

// eg:定义一个Person类,有name和age属性,有eat/sleep/work方法

type Person struct {
	name string
	age  int
}

func (per Person) eat() {
	fmt.Println("eat...")
}

func (per Person) sleep() {
	fmt.Println("sleep...")
}

func (per Person) work() {
	fmt.Println("work...")
}

func main() {
	per := Person{
		name: "klaus",
		age:  20,
	}
	fmt.Printf("per: %v\n", per)
	per.eat()
	per.sleep()
	per.work()

}
