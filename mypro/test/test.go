package main

import "fmt"

func main() {
	var name = "tom"
	var age = 20
	var b = true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)

}

//类型推断 (声明变量时，可以根据初始化值进行类型推导，从而省略类型)

// 批量声明 (使用一个var关键字，把一些变量写在一个括号()里)
/* eg: var identifier type
   var 声明变量关键字
   identifier 变量名称
   type 变量类型 */
/*package main
func main(){
	var(
		name string
		age int
		ok bool
	)
}

/*	// 批量初始化
	go语言在声明变量的时候，会自动为变量对应的内存区域进行初始化操作。每个变量会被初始化成
	其类型的默认值。eg：整型和浮点型变量默认值为0，字符串变量默认值为空字符串""".
	布尔型变量默认值为false。切片、函数、指针变量的默认值为nil.
	var name, age, b = "tom", 20, true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b) */

// 短变量声明 ：= (在函数内部，可以使用 :=运算符对变量进行声明和初始化) 注：函数外部不可使用！
/*name := "tom"
	age := 20
	b := true

	fmt.Printf("name: %v\n", name)
    fmt.Printf("age: %v\n", age)
    fmt.Printf("b: %v\n", b) */

//email :="pilw818@gmail.com"
//var email ="pilw818@gmail.com"

/* func getNameAndAge() (string, int) {
	return "tom", 20
} */

//如果我们接收到多个变量，有一些变量使用不到，可以使用下划线_表示变量名称，这种变量叫匿名变量

/* func main() {
	_, age := getNameAndAge()
	//fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
} */
