package main

import "fmt"

func main() {

	//go语言中if语句使用提示：
	/* 	1.不需要使用括号将条件包含起来
	   	2.大括号{}必须存在，即使只有一行语句
	   	3.左括号必须在if或else的同一行
	   	4.在if之后，条件语句之前，可以添加变量初始化语句，使用；进行分隔
		5.在有返回值的函数中，最终的return不能在条件语句中 */

	/* a := 1
	b := 2
	if a>b
	    fmt.Println("a")
	else
	    fmt.Println("b") //此写法错误，其他语言中一条语句可以不加大括号即用最外面的就行，golang中必须要加 */

	/* 	a := 1
	   	b := 2

	   	//gofmt
	   	if a > b {
	   		fmt.Printf("a: %v\n", a)
	   	} else {
	   		fmt.Printf("b: %v\n", b)
	   	} */

	//根据布尔值flag判断

	/* flag1 := true
	if flag1 {
		fmt.Printf("flag: %v\n", flag1)
	} */
	//根据年龄判断是否成年
	/* 	age := 20  //变量写在if语句外
	   	if age > 18 {
	   		fmt.Println("你已经成年")
	   	} else {
	   		fmt.Println("你还未成年")
	   	} */
	//写在里面
	if age := 17; age > 18 { //初始变量可以声明在布尔表达式中
		fmt.Println("你已经成年")
		fmt.Printf("age: %v\n", age) //注意它的作用域 在if语句里可打印
	} else {
		fmt.Println("你还未成年")
		fmt.Printf("age: %v\n", age)
	}
	//注意go语句中不能使用0或者非0表示真假

}
