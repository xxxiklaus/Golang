/*
	golang结构体作为函数参数

go结构体可以像普通变量一样，作为函数的参数，传递给函数，分两种情况
1.直接传递结构体，这是一个副本（拷贝），在函数内部不会改变外部结构体内容
2.传递结构体指针，这时在函数内部，能够改变外部结构体内容
*/
package main

import "fmt"

type Person struct {
	id   int
	name string
}

//值传递 拷贝了一份副本 （直接传递）
func showPerson(per Person) {
	per.id = 101
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

//传递结构体指针
func showPerson1(per *Person) {
	per.id = 102
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func main() {
	klaus := Person{
		id:   100,
		name: "klaus",
	}

	per := &klaus

	fmt.Printf("klaus: %v\n", klaus)

	showPerson1(per)

	fmt.Println("-----")
	fmt.Printf("per: %v\n", per)
	/* 	fmt.Printf("klaus: %v\n", klaus)
	   	fmt.Println("-------")
	   	showPerson(klaus)
	   	fmt.Printf("klaus: %v\n", klaus)
	*/
}
