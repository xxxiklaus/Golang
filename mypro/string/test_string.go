// gilang字符串
/* 一个Go语言字符串是一个任意字节的常量序列。 []byte
在Go语言中，字符串字面量使用双引号""或者反引号 ' 来创建。双引号用来创建可解析的字符串
支持转义，但不能用来引用多行；反引号用来创建原生的字符串字面量，可能由多行组成
但不支持转义，并且可以包含除了反引号外其他所有字符。双引号创建可解析的字符串应用最广泛
反引号用来创建原生的字符串则多用于书写多行消息，HTML以及正则表达式。 */
package main

import (
	"fmt"
	"strings"
)

func main() {

	//字符串函数

	s := "Hello World"
	fmt.Printf("len(s): %v\n", len(s))                                               //求字符串长度
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))               // 分割字符串  // +或fmt.Sprintf拼接字符串
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello")) // 判断是否包含该内容
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))                       // 小写
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))                       // 大写

	fmt.Printf("strings.HasPrefix(\"Hello\"): %v\n", strings.HasPrefix(s, "Hello")) // 判断前缀（是否大写）
	fmt.Printf("strings.HasSuffix(s): %v\n", strings.HasSuffix(s, "world"))         //  判断后缀（是否小写）

	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll")) // 判断子串出现的位置
	fmt.Printf("strings.LastIndex(s): %v\n", strings.LastIndex(s, "l"))  // 判断子串最后出现的位置

	//go语言字符串切片操作
	/* 	s := "hello world"
	   	a := 2
	   	b := 5

	   	fmt.Printf("s[a]: %c\n", s[a])     // %c只占位一个字符char
	   	fmt.Printf("s[a:b]: %v\n", s[a:b]) // %v代表任意类型
	   	fmt.Printf("s[a:]: %v\n", s[a:]) // 表示从2开始到最后
	   	fmt.Printf("s[:b]: %v\n", s[:b]) // 表示从0开始到当前b
	*/
	/* 	go语言字符串转义字符
	   	\r 回车符(返回行首)
	   	\n 换行符(直接跳到下一行同列位置)
	   	\t 制表符
	   	\' 单引号
	   	\"" 双引号
	   	\\ 反斜杠 */

	// s := "Hello\tworld"  //制表符 4个空格

	/* s := "c:\\program files\\a"
	   fmt.Printf("s: %v\n", s) */

	/* s := "I'm tom"
	   fmt.Printf("s: %v\n", s) */

	/* 	s := "hello \n world"
		fmt.Printf("s: %v\n", s)
	 	print(s + "\n")
		print(s)  */

	//字符串连接
	//虽然Go语言中的字符串是不可变的，但是字符串支持 + 级联操作和 += 追加操作

	/* 	var s string = "隋姐，明天几点开始种树？"
	   	var s1 = "你今天岂不是没看电影？"
	   	s3 := "为啥晚上不来种树了，对了那个青之芦苇更新了嘞"
	   	s4 := "累死了，今天就这样给隋姐汇报一下今天干了啥吧"
	   	s5 := `
	   	我今天写了6个test，平均每个test大概70行
	   	有俩个test100行，分别是const、datatype、
	   	bool、number、shortcuts以及现在这个string
	   	即语言常量、语言数据类型、语言布尔类型、
	   	语言数字类型、快捷键设置以及golang字符串
	   	看了8个视频的样子没有昨天多 但今天的东西难一点
	   	涉及到二进制、八进制和十六进制以及实数虚数还有个
	   	叫浮点类型的东西 报了不少错 所以慢了点 汇报完毕`

	   	fmt.Printf("s: %v\n", s)
	   	fmt.Printf("s1: %v\n", s1)
	   	fmt.Printf("s3: %v\n", s3)
	   	fmt.Printf("s4: %v\n", s4)
	   	fmt.Printf("s5: %v\n", s5) */

	/*
		var buffer bytes.Buffer
		buffer.WriteString("tom")
		buffer.WriteString(",")
		buffer.WriteString("20")

		fmt.Printf("buffer.String(): %v\n", buffer.String())

		这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符
		的长度，还可以用buffer.Grow()接口来设置capacity */

	/* 	name := "tom"
	   	age := "20"

	   	s := strings.Join([]string{name, age}, ",")
	   	fmt.Printf("s: %v\n", s)
	*/ //join会根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小
	//的内存，一个一个字符串填入，在已有一个数组的情况下 ，这种效率会很高，但是
	//本来没有，去构造这个数据的代价也不小

	/* 	s1 := "tom"
	   	s2 := "20"

	   	//内部使用{}byte实现，不像直接运算符这种会产生很多临时的字符串，但是内部的
		逻辑比较复杂，有很多额外的判断，还用到了interface，所以性能也不是很好

	   	msg := fmt.Sprintf("name=%s,age=%s", s1, s2)
	   	fmt.Printf("msg: %v\n", msg) */

	/* 	name := "tom"
	   	age := "20"
	   	msg := name + "" + age
	   	fmt.Printf("msg: %v\n", msg)
	   	fmt.Println("-------------")
	   	msg = " "
	   	msg += name
	   	msg += " "
	   	msg += age
	   	fmt.Printf("msg: %v\n", msg) */

	/* golang里面的字符串都是不可变的，每次运算都会产生一个新的字符串，所以会产生
	   很多临时的无用的字符串，不仅没用，还会给gc带来额外的负担，所以性能比较差 */

}
