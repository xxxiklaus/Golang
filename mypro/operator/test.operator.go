package main

import "fmt"

func getA() int {
	return 100
}

// go语言内置运算符

func main() {
	//算术运算符
	/* 	a := 100
	   	b := 20 */
	/*    	fmt.Printf("(a + b): %v\n", (a + b)) //加
	fmt.Printf("(a - b): %v\n", (a - b)) //减
	fmt.Printf("(a / b): %v\n", (a / b)) //乘
	fmt.Printf("(a * b): %v\n", (a * b)) //除
	x := a % b
	fmt.Printf("x: %v\n", x) //求余  */

	//注: ++(自增)和--(自减)在Go语言中是单独的语句！并不是运算符
	/* 	a := 100
	   	a++
	   	fmt.Printf("a: %v\n", a)
	   	a--
	   	fmt.Printf("a: %v\n", a) */

	//关系运算符
	/* 	fmt.Printf("(a == b): %v\n", (a == b)) //检查两个值是否相等，相等返回true，否之返回false
	   	fmt.Printf("(a != b): %v\n", (a != b)) //检查两个值是否不相等，不相等返回true，否之返回false
	   	fmt.Printf("(a > b): %v\n", (a > b)) //检查左边值是否大于右边值，是返回true，否之返回false
	   	fmt.Printf("(a >= b): %v\n", (a >= b)) //检查左边值是否大于等于右边值，是返回true，否之返回false
	   	fmt.Printf("(a < b): %v\n", (a < b)) //检查左边值是否小于右边值，是返回true，否之返回false
	   	fmt.Printf("(a <= b): %v\n", (a <= b)) //检查左边值值是否小于等于右边值，是返回true，否之返回false */

	//逻辑运算符
	/* 	a := true
	   	b := false

	   	r := a && b  // 逻辑运算AND符，两边都是true，则为true，否之则为false
	   	fmt.Printf("r: %v\n", r)
	   	r = a || b //逻辑OR运算符 两边有一个为true，则为true，否之则为false
	   	fmt.Printf("r: %v\n", r)
	   	fmt.Printf("a: %v\n", !a) //逻辑NOT运算符，如果条件为true，则为false，否则为true */

	//位运算 位运算符对整数在内存中的二进制进行操作
	/* 	a := 4
	   	fmt.Printf("a: %b\n", a) //0100
	   	b := 8
	   	fmt.Printf("b: %b\n", b) //1000

	   	r := a & b //参与运算的两数各对应的二进位相与（两位均为1才为1）
	   	fmt.Printf("r: %v\n", r)
	   	r = a | b //参与运算的两数各对应的二进位相或 （两位有一个为1就为1）
	   	fmt.Printf("r: %v\n", r)
	   	r = a ^ b //参与运算的两数各对应的二进位相异或（两个不一样则为1）
	   	fmt.Printf("r: %v\n", r)
	   	fmt.Printf("(a << 2): %v\n", (a << 2)) //左移n位就是乘以2的n次方
	   	//"a<<b"是把a的各二进位全部左移b位，高位丢弃，低位补0
	   	fmt.Printf("(a >> 2): %v\n", (a >> 2))//右移n位就是除以2的n次方
	   	//"a>>b"是把a的各二进位全部右移b位 */

	//赋值运算符
	a := 100
	a = 200 // = 简单的赋值运算符，将一个表达式的值赋给一个左值
	fmt.Printf("a: %v\n", a)
	a = getA() //100 函数返回赋值
	fmt.Printf("a: %v\n", a)
	a += 100 // a = a + 100 //相加后再赋值
	fmt.Printf("a: %v\n", a)
	a -= 100 //相减后赋值
	fmt.Printf("a: %v\n", a)
	a *= 100 //相乘后赋值
	fmt.Printf("a: %v\n", a)
	a /= 100 //相除后赋值
	fmt.Printf("a: %v\n", a)
	a %= 100 //求余后再赋值
	fmt.Printf("a: %v\n", a)
	a &= 100 //按位与后赋值
	fmt.Printf("a: %v\n", a)
	a |= 100 //按位或后赋值
	fmt.Printf("a: %v\n", a)
	a ^= 100 //按位异或后赋值
	fmt.Printf("a: %v\n", a)

}
