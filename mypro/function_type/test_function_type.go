/* golang函数类型与函数变量
可以使用type关键字来定义一个函数变量，语法格式：
type fun func(int, int)int
上面语句定义了一个fun函数类型，它是一种函数类型，这种函数类型接收
两个int类型的参数并且返回一个int类型的返回值 */

package main

import "fmt"

//下面我们定义这样结构的两个函数，一个求和，一个比较大小

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	type f1 func(int, int) int
	var ff f1
	ff = sum
	r := ff(88, 99)
	fmt.Printf("r: %v\n", r)

	ff = max
	r = ff(1557, 4396)
	fmt.Printf("r: %v\n", r)

}
