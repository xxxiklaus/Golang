/*
	golang类型定义和类型别名

类型定义的语法 type NewType Type  // 关键字 新定义的类型 原始类型
go语言类型别名 类型别名语法： type NewType = Type
go语言类型定义和类型别名的区别
1.类型定义相当于定义了一个全新的类型，与之前的类型不同;但是类型别名并没有
定义一个新的类型，而是使用一个别名来替换之前的类型
2.类型别名只会在代码中存在，在编译完成之后并不会存在该别名
3.因为类型别名和原来的类型是一致的，所以原来类型所拥有的方法，类型别名中
也可以调用，但是如果是重新定义的一个类型，那么不可以调用之前的任何方法
*/
package main

import "fmt"

func main() {
	/* 	type MyInt int //类型定义
	   	var i MyInt
	   	i = 100
	   	fmt.Printf("i: %v i:%T\n", i, i)
	*/
	type MyInt1 = int //类型别名定义
	var i MyInt1      //i其实还是int类型
	i = 100
	fmt.Printf("i: %v i:%T\n", i, i)
}
