/*
	golang方法接收者类型

结构体实例，有值类型和指针类型，那么方法的接收者是结构体，那么也有值类型
和指针类型.区别就是接收者是否复制结构体副本.值类型复制，指针类型不复制
*/
package main

import "fmt"

func test1() {

	p1 := Person{
		name: "klaus",
	}

	p2 := &Person{
		name: "klaus",
	}

	fmt.Printf("p1: %T\n", p1)  //值类型
	fmt.Printf("p2: %T+\n", p2) //指针类型

}

type Person struct {
	name string
}

func showPerson1(per Person) {
	per.name = "klaus..."

}

func showPerson2(per *Person) {
	//per.name自动解引用
	//(*per).name ="klaus..."
	per.name = "klaus..."

}

func (per Person) showPerson3() {
	per.name = "klaus..."

}

func (per *Person) showPerson4() {
	//per.name自动解引用
	//(*per).name ="klaus..."
	per.name = "klaus..."

}

func main() {
	p1 := Person{
		name: "klaus",
	}

	p2 := &Person{
		name: "klaus",
	}
	/* showPerson1(p1)
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("------")
	showPerson2(p2)
	fmt.Printf("p2: %v\n", p2) */

	p1.showPerson3()
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("------")
	p2.showPerson4()

	fmt.Printf("p2: %v\n", *p2)

}
