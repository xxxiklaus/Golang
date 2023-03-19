// golang函数的返回值
/* 函数可以有0或多个返回值，返回值需要指定数据类型，返回值通过return关键字来指定
return可以有参数，也可以没有参数，这些返回值可以有名称，也可以没有名称
go中的函数可以有多个返回值
1.return关键字指定了参数时，返回值不可以用名称.如果return省略参数，则返回值部分必须带名称
2.当返回值有名称时，必须使用括号包围，逗号分隔，即使只有一个返回值
3.但即使返回值命名了，return中也可以强制指定其他返回值的名称，也就是说return的优先级更高
4.命名的返回值是预先声明好的，在函数内部可以直接使用，无需再次声明，
命名返回值的名称不能和函数参数名称相同，否则报错提示变量重复定义
5.return中可以有表达式。但不能出现赋值表达式，这和其它语言可能有所不同，
eg：return a + b 是正确的，但是 return c = a + b是错误的 */
package main

import (
	"fmt"
)

// 没有返回值
func test1() {
	fmt.Println("我没有返回值，只是进行一些计算")
}

//有一个返回值

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret

}

// 多个返回值,且在return中指定返回的内容
func test2() (name string, age int) {
	name = "klaus"
	age = 22
	return name, age
	// return 等价于 前者 （多个返回值返回值名称没有被使用）

}

//return覆盖命名返回值，返回值名称没有被使用

func test3() (name string, age int, email string) {
	n := "klaus"
	a := 22
	e := "pilw818@gmail.com"
	return n, a, e
}

/* Go中经常会使用其中一个返回值作为函数是否执行成功，是否有错误信息的判断条件
eg: return value,exists、return value,ok、return value,err等
当函数的返回值过多时，例如有4个以上的返回值，应该将这些返回值收集到容器中
然后以返回容器的方式去返回eg:同类型的返回值可以放进slice中，不同类型的
返回值可以放进map中
当函数有多个返回值时，如果其中某个或某几个返回值不想使用，可以通过下划线_
来丢弃这些返回值.eg：下面的test4函数的俩个返回值，调用该函数时，丢弃了
第二个返回值b，只保留了第一个返回值a赋值给了变量a. */

func test4() (int, int) {
	return 1, 2
}

func main() {
	// test1()
	// r := sum(4, 6)
	// fmt.Printf("r: %v\n", r)
	/* 	name, age := test2()
	   	fmt.Printf("name: %v\n", name)
	   	fmt.Printf("age: %v\n", age) */
	/* n, a, e := test3()
	fmt.Printf("n: %v\n", n)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("e: %v\n", e) */
	_, x := test4()
	fmt.Printf("x:%v\n", x)

}
