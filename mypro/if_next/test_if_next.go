//go语言中if嵌套语法
/* if 布尔表达式1{
	1为true时执行
	 if 布尔表达式2{
		2为true时执行
	 }
}*/

package main

import "fmt"

// 判断3个数的大小
//a>b a>c a
// b>a b>c else c
func f1() {
	a := 100
	b := 200
	c := 3
	if a > b {
		if a > c {
			fmt.Println("a")
		}
	} else {
		// b>a
		if b > c {
			fmt.Println("b")
		} else {
			fmt.Println("c")
		}
	}
}

//判断男生还是女生，是否成年
func f2() {
	gender := "女生"
	age := 16
	if gender == "男生" {
		fmt.Println("男生")
		if age > 18 {
			fmt.Println("成年")
		} else {
			fmt.Println("未成年")
		}
	} else {
		fmt.Println("女生")
		if age > 18 {
			fmt.Println("成年")
		} else {
			fmt.Println("未成年")
		}
	}
}
func main() {
	// f1()
	f2()

}
