//go语言中if else if语法
/* if 布尔表达式1{
	do something
}elseif 布尔表达式2{
	do something else
} else{
	catch-all or default
}*/

package main

import "fmt"

//根据分散判断等级
func f5() {
	score := 90

	if score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score >= 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

//输入星期几的第一个字母来判断一下是星期几，如果第一个字母一样则继续判断第二个字母
// Monday Tuesday Wednesday Thursday Friday Saturday Sunday
func f6() {
	var c string
	fmt.Println("输入一个字符")
	fmt.Scan(&c)

	if c == "S" {
		fmt.Println("第二个字符,")
		fmt.Scan(&c)
		if c == "a" {
			fmt.Println("Saturday")
		} else if c == "u" {
			fmt.Println("Sunday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "F" {
		fmt.Println("Friday")
	} else if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("输入第二个字符")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuesday")
		} else if c == "h" {
			fmt.Println("Thursday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	}

}

func main() {
	f6()
	// f5()

}
