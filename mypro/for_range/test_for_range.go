// golang for range循环
/* go语言中可以使用for range遍历数组、切片、字符串、map及通道（channel）
通过for range遍历的返回值有以下规律：
1.数组、切片、字符串返回索引和值
2.map返回键和值
3.通道(channel)返回通道内的值 */
package main

import "fmt"

/* func f1() {
	/* 	var a = [...]int{1, 2, 3}
	   	for i, v := range a {
	   		fmt.Printf("i: %v\n", i)
	   		fmt.Printf("v: %v\n", v)

	   	}
} */

/* func f2() {
	var a = [...]int{1, 2, 3}
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
} */

/* func f3() {
	var s = []int{1, 2, 3}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
} */

/* func f4() {
	//key : value

	var m = make(map[string]string, 0)
	m["name"] = "klaus"
	m["age"] = "20"
	m["email"] = "pilw818@gmail.com"

	for key, value := range m {
		fmt.Printf("%v:%v\n", key, value)
	}
} */

//字符串
func f5() {
	s := "hello go"
	for _, v := range s {
		fmt.Printf("v: %c\n", v)
	}
}
func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	f5()

}
