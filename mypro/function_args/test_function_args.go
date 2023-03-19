/*
	golang函数的参数

go语言函数可以有0或多个参数，参数需要指定的数据类型
声明函数时的参数列表叫做形参，调用时传递的函数叫做实参
go语言是通过传值的方式传参的，意味着传递给函数的是拷贝后的副本，
所以函数内部访问、修改的也是这个副本
go语言可以使用变长参数，有时候并不能确定参数的个数，可以使用变长
参数，可以在函数定义语句的参数部分使用ARGS...TYPE的方式，这时会
将...代表的参数全部保存到一个名为ARGS的slice中，注意这些参数的数据类型都是TYPE
*/
package main

import "fmt"

/* func sum(a int, b int) int { //形参列表
	return a + b
} */

//演示参数传递，按值传递

/* func f1(a int) { //copy
	a = 100
	fmt.Printf("a: %v\n", a)
} */

func f2(s []int) {
	s[0] = 1000
}

//变长参数

func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)

	}
}

func f4(name string, ok bool, args ...int) { //前边是固定参数 后边是可变参数
	fmt.Printf("name: %v\n", name)
	fmt.Println("-------")
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("-------")
	for _, v := range args {
		fmt.Printf("v: %v\n", v)

	}
}

func main() {
	/*   r := sum(1,2)  //实参列表
	     fmt.Printf("r: %v\n", r) */
	/* 	x := 200
	   	f1(x)
	   	fmt.Printf("x: %v\n", x)
	*/
	//map、slice、interface、channel这些数据类型本身就是"指针"类型的，所以
	// 就算是拷贝传值也是拷贝的指针，拷贝后的参数仍然指向底层数据结构，所以修改
	// 它们可能会影响外部数据结构的值

	/* s := []int{1, 2, 3}
	f2(s)
	fmt.Printf("s: %v\n", s) */
	/* 	f3(1, 2, 3, 4)
	   	f3(3, 4, 5, 6, 7) */
	f4("klaus", true, 1, 2, 3, 4)

}
