/*
	golang继承

golang本质上没有OOP概念,也没有继承概念,但是可以通过结构体嵌套实现这个特性
*/
package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Tiger struct {
	Animal //可以理解为继承
	color  string
}

type Hawk struct {
	Animal
	kind string
}

func main() {
	tiger := Tiger{
		Animal: Animal{name: "东北虎", age: 3},
		color:  "yellow and white",
	}
	hawk := Hawk{
		Animal: Animal{name: "猫头鹰", age: 1},
		kind:   "owl",
	}
	tiger.eat()
	tiger.sleep()
	fmt.Printf("tiger.name: %v\n", tiger.name)
	fmt.Printf("tiger.age: %v\n", tiger.age)
	fmt.Printf("tiger.color: %v\n", tiger.color)

	hawk.eat()
	hawk.sleep()
	fmt.Printf("hawk.name: %v\n", hawk.name)
	fmt.Printf("hawk.age: %v\n", hawk.age)
	fmt.Printf("hawk.kind: %v\n", hawk.kind)

}
