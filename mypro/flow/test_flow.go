// go语言中的流程控制
package main

import "fmt"

func main() {
	//顺序语句
	fmt.Println("aaa")
	fmt.Println("bbb")

	//条件语句
	/* 	条件语句是用来判断给定的条件是否满足（表达式值是否为true或者false）
	   	并根据判断的结果（真或假）决定执行的语句 */
	a := 100
	//if语句 由一个布尔表达式后缀紧跟一个或多个语句组成
	//if...else语句 if语句后可以使用可选的else语句，else语句中的表达式
	//在布尔表达式为false时执行
	//if嵌套语句 你可以在if或else if中嵌入一个或多个if或else if语句
	if a > 20 {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}

	//switch语句 switch语句用于基于不同条件执行不同动作
	//select语句 select语句类似于switch语句，但是select语句会随机执行
	//一个可运行的case，如果没有case可运行，它将阻塞，直到有case可运行
	switch a {
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("default")
	}

	//go语言中的循环语句
	//go语言中的循环只有for循环，去除了while、do while循环，使用起来更简洁

	//for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	//for range循环
	x := [...]int{1, 2, 3}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}
}
