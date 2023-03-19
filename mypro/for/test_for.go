//golang for循环语句
// go语言中的for循环,只有for关键字，去除了像其他语言中的while和do while
/* go语言for循环语法
for 初始语句;条件表达式;结束语句{
	循环体语句
} */
package main

import "fmt"

func main() {
	//循环输出1到10
	/* for i := 1; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	} */
	//初始条件可以写到外面
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
	//初始条件和结束条件都可以省略
	/* 	i := 1 //初始条件
	   	for i <= 10 {
	   		fmt.Printf("i: %v\n", i)
	   		i++ //结束条件
	   	} */ //这种情况类似其他语言中的while循环
	//永真循环
	/* for {
		fmt.Println("我一直在执行~")
	} */
	// for循环可以通过break、goto、return、panic语句强制退出循环
}
