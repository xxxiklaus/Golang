/*
	golang标准库log

golang内置了log包,实现简单的日志服务.通过调用log包的函数
可以实现简单日志打印功能
函数系列        作用
print        单纯打印日志
panic        打印日志,抛出panic异常
fatal        打印日志,强制结束程序(os.Exit(l)),defer函数不会执行

log配置
标志log配置
默认情况下log只会打印除时间,但是实际情况下我们可能还需要获取文件名
行号等信息,log包提供给我们定制的接口.log包提供两个标准log配置相关方法:
func Flags() int //返回标准log输出配置
func SetFlags(flag int) //设置标准log输出配置

flag参数
const (

	/控制输出日志信息的细节,不能控制输出的顺序和格式.输出日志在每一项后会有一个冒号分隔
	/eg: 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message

	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger

)
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func test1() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "klaus"
	age := 22
	log.Println(name, " ", age)

}

func test2() {
	defer fmt.Println("panic 结束后再执行")
	log.Print("my log")
	log.Panic("my panic")
	fmt.Println("end...")

}

func test3() {
	defer fmt.Println("defer...")
	log.Print("my log")
	log.Fatal("fatal...")
	// os.Exit(1)
	fmt.Println("end...")
}

// 标准日志设置示例
func test4() {
	i := log.Flags()
	//二进制表示 Ldate=01 Ltime=10 Ldate|Ltime=3
	fmt.Printf("i: %v\n", i) //i=3

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("my log...")
}

// 日志前缀示例
// func Prefix() string /返回日志的前缀配置
// func SetPrefix(prefix string) //设置日志前缀
func test5() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("Mylog:")
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// defer f.Close()
	if err != nil {
		log.Fatal("日志文件错误")
	}
	log.SetOutput(f)

	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.Print("my log...")

}

//自定义logger
//log包为我们提供了内置函数,让我们能自定义logger.从效果上来看,就是将标准日志配置
//日志前缀配置,日志输出位置整合到一个函数中,使日志配置不那么繁琐.
//func New(out io.Writer, prefix string, flag int) *Logger 实现自定义logger

func test6() {
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("日志文件错误")
	}

	logger := log.New(f, "Mylog:", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Print("my log...")

}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	test6()

}
