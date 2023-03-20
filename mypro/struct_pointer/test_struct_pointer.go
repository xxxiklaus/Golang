/*
	golang 结构体指针

结构体指针和普通变量指针相同
*/
package main

import "fmt"

func test1() {
	type Person struct {
		id    int
		age   int
		name  string
		email string
	}
	klaus := Person{
		id:    101,
		age:   22,
		name:  "klaus",
		email: "pilw818@gmail.com",
	}
	var p_person *Person
	p_person = &klaus

	fmt.Printf("klaus: %v\n", klaus)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)

}

func main() {
	/* var name string
	name = "klaus"
	var p_name *string //*p_name 指针类型
	p_name = &name     // &name 取name地址
	fmt.Printf("name: %v\n", name)
	fmt.Printf("p_name: %v\n", p_name)   // 输出指针地址
	fmt.Printf("*p_name: %v\n", *p_name) //输出指针指向的内容量 */
	// test1()
	//使用new关键字创建结构体指针
	type Person struct {
		id   int
		name string
	}

	var p_person = new(Person)
	fmt.Printf("p_person: %T\n", p_person)

}
