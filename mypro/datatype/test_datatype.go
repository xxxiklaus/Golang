/*
	go语言数据类型

数据类型用于声明函数和变量
数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候
才需要申请大内存，这样就可以充分利用内存
Go语言按类别有以下几种数据类型
1布尔型： 布尔型的值只可以是常量true或者false，eg： var b bool = true.
2数字类型：整型int和浮点型float32、float64，Go语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码.
3字符串类型：字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的.
Go语言的字符串字节使用UTF-8编码标识Unicode文本.
4派生类型：包括：(a)指针类型(Pointer) (b)数组类型 (c)结构化类型(struct) (d)Channel类型
(e)函数类型 (f)切片类型 (g)接口类型(interface) (h)Map类型
*/
package main

import "fmt"

/*  func a() {

} //函数类型  */

func main() {

	// fmt.Printf("%T\n", a)
	var name string = "tom"
	age := 20
	b := true
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b) //打印类型 */

	/* 	a := 100
	   	p := &a
	   	fmt.Printf("%T\n", a)
	   	fmt.Printf("%T\n", p) */ // 指针类型

	/* 	a := [2]int{1, 2}
	   	fmt.Printf("%T\n", a) */ //数组类型
	/* 	a := []int{1, 2, 3}
	   	fmt.Printf("%T\n", a) */ //切片类型

}

//数字类型
/* Go基于架构的类型，eg：int、uint、uintptr
1：uint8 无符号8位整型(0到255)
2：uint16无符号16位整型(0到65535)
3：uint32无符号32位整型(0到4294967295)
4：uint64无符号64位整型(0到18446744073709551615)
5：int8有符号8位整型(-128到127)
6：int16有符号16位整型(-32768到32767)
7：int32有符号32位整型(-2147483648到2147483647)
8：int64有符号64位整型(-9223372036854775808到9223372036854775807)
*/
//浮点型
/* 1：float32 IEEE -754 32位浮点型数
2：float64 IEEE -754 64位浮点型数
3：complex64 32位实数和虚数
4：complex128 64位实数和虚数 */
// 其他数字类型
/* 1： byte和uint8类似
2： rune类似int32
3： uint32或64位
4：int与uint 一样大小
5： uintptr 无符号整型，用于存放一个指针 */
