// golang高阶函数
package main

import "fmt"

//golang函数作为参数
func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

//golang函数作为返回值

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	// test("golang", sayHello)
	ff := cal("+")
	r := ff(1557, 4396)
	fmt.Printf("r: %v\n", r)
	fmt.Println("------")
	ff = cal("-")
	r = ff(4396, 1557)
	fmt.Printf("r: %v\n", r)
}
