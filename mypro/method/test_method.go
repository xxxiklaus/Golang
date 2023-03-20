/*
	golang方法

go语言没有面向对象特性，也没有类对象的概念，但是，可以使用结构体来模拟
这些特性，我们都知道面向对象里面有类方法等概念.我们也可以声明一些方法属于某个结构体
golang方法语法
Go中的方法，是一种特殊函数，定义于struct之上（与struct关联、绑定），被称为
struct的接收者receiver，通俗的讲，方法就是接收者的函数，语法格式如下：
type mytype struct {}

func (recv mytype)my_method(para) return_type{}
func (recv *mytype)my_method(para) return_type{}
mytype:定义一个结构体
recv:接受该方法的结构体（receiver）
my_method:方法名称
para:参数列表
return_type:返回值类型
*/
package main

import "fmt"

type Person struct {
	name string
}

//属性和方法分开来写

//(per Person)接收者receiver

func (per Person) eat() {
	fmt.Printf("%v,eat...\n", per.name)
}

func (per Person) sleep() {
	fmt.Printf("%v,sleep...", per.name)
}

type Customer struct {
	name string
	pwd  string
}

func (customer Customer) login(name string, pwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	fmt.Printf("customer.pwd: %v\n", customer.pwd)
	if name == "klaus" && pwd == "818" {
		return true
	} else {
		return false
	}

}

func main() {
	cus := Customer{
		name: "klaus",
		pwd:  "818",
	}

	b := cus.login("klaus", "818")
	fmt.Printf("b: %v\n", b)

	/* per := Person{
		name: "klaus",
	}
	per.eat()
	per.sleep() */

}

/* go语言方法的注意事项

1.方法的receiver type 并非一定要是struct类型，type类型别名、slice、map、channel、func类型等都可以
2.struct结合它的方法就等价于面向对象中的类。只不过struct可以和它的方法分开
并非一定要属于同一个文件，但必须属于同一个包
3.方法有两种接收类型：（T Type）和（T *Type），它们之间有区别
4.方法就是函数，所以Go中没有方法重载（overload）的说法，也就是说
同一个类型中的所有方法名必须都唯一
5.如果receiver是一个指针类型，则会自动解除引用
6.方法和type是分开的，意味着实例的行为（behavior）和数据存储（field）是分开的，但它们通过receiver建立起关联关系
*/
