/*
	golang指向数组的指针

语法 var ptr [MAX]*int; 表示数组里面的元素的类型是指针类型
*/
package main

import "fmt"

func main() {

	a := [3]int{1, 2, 3}
	var pa [3]*int

	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]

	}
	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(pa); i++ {
		fmt.Printf("%v\n", *pa[i])

	}

}
