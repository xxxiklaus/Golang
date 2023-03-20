/*
	golang通过接口实现OCP设计原则

面向对象的可复用设计的第一块基石,便是所谓的"开-闭"原则(Open-Closed-Principle,常缩写为OCP)
虽然golang不是面向对象语言,但是也可以模拟实现这个原则.对扩展是开放的,对修改是关闭的(已经写好的代码不可修改)
*/
package main

import (
	"fmt"
)

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

// Dog 实现Pet接口
func (dog Dog) eat() {

	fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

type Cat struct {
	name string
	age  int
}

// Cat实现Pet接口
func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}

type Person struct {
	name string
}

// pet 既可以传递Dog也可以传递Cat
func (per Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}
func main() {

	dog := Dog{
		name: "大黄",
		age:  2,
	}
	fmt.Printf("dog.name: %v\n", dog.name)
	fmt.Printf("dog.age: %v\n", dog.age)
	cat := Cat{
		name: "小黑",
		age:  1,
	}
	fmt.Printf("cat.name: %v\n", cat.name)
	fmt.Printf("cat.age: %v\n", cat.age)
	per := Person{
		name: "klaus",
	}
	fmt.Printf("per.name: %v\n", per.name)
	per.care(dog)
	per.care(cat)

}
