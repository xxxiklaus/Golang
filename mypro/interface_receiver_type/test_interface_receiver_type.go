/*
	golang接口值类型接收者和指针类型接收者

本质上和方法的值类型接收者和指针类型接收者的思考方法是一样的
值接收者是一个拷贝,是一个副本,而指针接收者,传递的是指针
*/
package main

import (
	"fmt"
)

type Pet interface {
	eat(string) string
}

type Dog struct {
	name string
}

/* func (dog Dog) eat() {
	fmt.Printf("dog: %p\n", &dog)
	fmt.Println("dog eat...")
	dog.name = "大黄"
} */

func (dog *Dog) eat(name string) string { //将Pet接口修改为指针接收者

	dog.name = "大黄ovo"
	fmt.Printf("dog: %p\n", dog)
	fmt.Printf("name: %v\n", name)
	return "吃得很好"
}

func main() {
	/* dog := Dog{
		name: "大黄",
	}
	fmt.Printf("dog: %p\n", &dog) 实现Pet接口(接收者是值类型)
	dog.eat()
	fmt.Printf("dog: %v\n", dog)
	从运行结果看,dog地址变了,说明是复制了一份,dog的name没变
	说明外面的变量没有被改变*/
	dog := &Dog{
		name: "大黄",
	}

	fmt.Printf("dog: %p\n", dog)
	s := dog.eat("火鸡")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("dog: %v\n", dog)

}

/* 正常情况下,当eat函数出现多次就会报错,提示会说函数已经声明过了
但interface后,同一个函数名在每个结构体作为接收者时,有不同的实现
类似于重载(overload)或者多态 */
