//golang流程控制关键字break
// break语句可以结束for、switch、select的代码块
// go语言使用break注意事项
/* 1.单独在select中使用break和不使用break没有啥区别
2.单独在表达式switch语句，并且没有fallthrough，使用break和不使用break没有啥区别
3.单独在表达式switch语句，并且有fallthrough，使用break能够终止fallthrough后面的case语句的执行
4.带标签的break，可以跳出多层select/switch作用域。让break更加灵活，写法更加简单灵活
不需使用控制变量一层一层跳出循环，并没有带break的只能跳出当前语句块。 */
package main

import "fmt"

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break //跳出循环
		}
	}
}

func test2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
		// break
	case 2:
		fmt.Println("2")
		// break 跳出switch里的fallthrough
		fallthrough
	case 3:
		fmt.Println("3")

	}
}

func test3() {
MY_LABEL: // 标签的定义即标签的名称后加 ：

	for i := 0; i < 10; i++ {
		if i >= 5 {
			break MY_LABEL //跳转到标签处
		}
		fmt.Printf("i: %v\n", i)
	}
	fmt.Println("END...")
}
func main() {

	// test1()
	// test2()
	test3()
}
