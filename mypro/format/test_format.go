package main

import "fmt"

//golang格式化输出
type WebSite struct {
	Name string
}

func main() {

	// %v 代表var变量的意思
	// fmt.Printf("\"hello\": %v\n", "hello")

	// 定义结构体变量

	site := WebSite{Name: "duoke360"}

	// 占位符
	//普通占位符

	fmt.Printf("site: %v\n", site)  //%v  相应值的默认格式  输出：site: {duoke360}
	fmt.Printf("site: %+v\n", site) //%+v 在打印结构体时"加号"标记"%+v"会添加字段名  输出：site: {Name:duoke360}
	fmt.Printf("site: %#v\n", site) // %#v 相应值的go语法表示 输出：site: main.WebSite{Name:"duoke360"}
	fmt.Printf("site: %T\n", site)  // %T 相应值的类型的go语法表示 输出：site:main.WebSite
	a := 100
	fmt.Printf("a: %T\n", a) // 输出a为int类型
	fmt.Println(" %%")       // %% 字面上的百分号，并非值的占位符 输出： %%

	// 布尔类型占位符

	/* b := true
	   fmt.Printf("b: %t\n", b)  // %t 单词 true或者false */

	//整型类型占位符
	/* 	i := 8
	   	fmt.Printf("i: %v\n", i) （8）
	   	fmt.Printf("i: %b\n", i) // 二进制表示 （1000）
	   	i = 96
	   	fmt.Printf("i: %c\n", 0x4E2D) // 相应的Unicode码点所表示的字符 （中）
	   	fmt.Printf("i: %d\n", i)      //十进制表示 （96）
	   	fmt.Printf("i: %o\n", 10)     //八进制表示 （12）
	   	fmt.Printf("i: %q\n", 0x4E2D) //单引号围绕的字符字面值，由Go语法安全的转义 （'中'）
	   	fmt.Printf("i: %x\n", 100)    // 十六进制表示， 字母形式为小写的 a-f （64）
	   	fmt.Printf("i: %x\n", 1234)   （4d2）
	   	fmt.Printf("i: %X\n", 13)     //十六进制表示，字母形式为大写的A-F （D）
	   	fmt.Printf("i: %U\n", 0x4E2D) //Unicode格式，U+1234，等同于U+%04X  （U+4E2D）  */

	//浮点数和复数的组成部分(实部和虚部)
	// %b 无小数部分的，指数为2的幂的科学计数法，与strconv.FormatFloat的'b'转换格式一致

	/* 	fmt.Printf("i: %e\n", 10.2)     //科学计数法，eg：-123456p-78  1.020000e+01
	   	fmt.Printf("i: %E\n", 10.2)     //科学计数法，eg：-123456p-78  1.020000E+01
	   	fmt.Printf("i: %f\n", 10.2)     //有小数点而无指数 eg：123.456   10.200000
	   	fmt.Printf("i: %g\n", 10.20)    //根据情况选择%e或%f以产生更紧凑的（无末尾的0）10.2
	   	fmt.Printf("i: %G\n", 10.20+2i) //根据情况选择%E或%f以产生更紧凑的（无末尾的0  (10.2+2i)
	*/

	//字符串与字节切片
	/* 	fmt.Printf("%s\n", []byte("github.com")) //输出字符串表示(string类型或[]byte)
	   	fmt.Printf("%q\n", "github.com")         //双引号围绕的字符串，由Go语法安全转义
	   	fmt.Printf("%x\n", "golang")             //十六进制，小写字母，每字节两个字符
	   	fmt.Printf("%X\n", "golang")             //十六进制，大写字母，每字节两个字符 */

	//指针

	/* 	x := 100
	   	p := &x
	   	fmt.Printf("i: %p\n", p) //%p 十六进制表示前缀0x 0xc000016098 */
}
