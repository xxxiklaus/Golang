//golang函数简介
// 函数是go语言中的“一级公民”，我们把所有的功能单元都定义在函数中，可以重复使用
// 函数包含函数的名称、参数列表和返回值类型，这些构成了函数的签名（signature）
// golang中函数特性
/* 1.go语言中有三种函数：普通函数、匿名函数（没有名称的函数）、方法定义在struct上的函数;
2.go语言中不允许函数重载（overload），也就是说不允许函数同名;
3.go语言中的函数不能嵌套函数，但可以嵌套匿名函数;
4.函数是一个值，可以将函数赋值给变量，使得这个变量也成为函数;
5.函数可以作为参数传递给另一个函数;
6.函数的返回值可以是一个函数;
7.函数调用的时候，如果有参数传递给函数，则先拷贝函数的副本，再将副本传递给函数;
8.函数参数可以没有名称.
*/
// golang中函数的定义和调用
// 函数在使用之前必须先定义，可以调用函数来完成某个任务。函数可以重复使用，从而达到代码重用
// golang函数定义语法
/* func function_name([parameter list])[return_types] {
	函数体
}
func: 函数由func 开始声明
function_name: 函数名称，函数名和参数列表一起构成函数签名
[parameter_list]: 参数列表，参数就像一个占位符，当函数被调用时，你可以直接传递给参数
这个值被称为实际参数.参数列表指定的是参数类型，顺序，参数个数.参数是可选的，也就是说
函数也可以不包含参数
return_types: 返回类型，函数返回一列值。return_types是该列值的数据类型。有些功能
不需要返回值，这种情况下return_types不是必须的.
函数体:函数定义的代码集合 */
package main

import "fmt"

//定义一个求和函数
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

//定义一个比较两个数大小的函数
func compare(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func main() {
	/* 	r := sum(1, 2)
	   	fmt.Printf("r: %v\n", r) */
	r := compare(1, 2)
	fmt.Printf("r: %v\n", r)
}
