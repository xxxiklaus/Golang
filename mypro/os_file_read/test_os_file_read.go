/*
	golang标准库os模块-File文件读操作

下面介绍File结构体相关的文件读操作
*/
package main

import (
	"fmt"
	"os"
)

// 文件打开关闭
func OpenClose() {
	//只能读
	/* 	f, err := os.Open("a.txt")
	   	if err != nil {
	   		fmt.Printf("err: %v\n", err)
	   	} else {
	   		fmt.Printf("f.Name(): %v\n", f.Name())
	   		f.Close()
	   	} */

	// 可以读写或者创建
	f, err := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}

}

// 创建文件
func createFile() {
	// 等价于, OpenFile(name,O_RDWR|O_CREATE|O_TRUNC,0666)
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	//第1个参数 目录默认:Temp 第二个参数 文件名前缀
	f1, _ := os.CreateTemp("", "temp")
	fmt.Printf("f1.Name(): %v\n", f1.Name())

}

//读操作 read

func readOps() {
	//循环读取
	/* 	f, _ := os.Open("a.txt")
	   	for {
	   		buf := make([]byte, 3)
	   		n, err := f.Read(buf)
	   		if err == io.EOF {
	   			break

	   		}
	   		fmt.Printf("n: %v\n", n)
	   		fmt.Printf("string(buf): %v\n", string(buf))
	   	}
	   	f.Close() */

	/* f, _ := os.Open("a.txt")
	buf := make([]byte, 4)
	n, _ := f.ReadAt(buf, 3) //从3开始读 4个字节
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf)) */

	//测试a目录下有b目录和c目录
	/* 	de, _ := os.ReadDir("packagetest")
	   	for _, v := range de {
	   		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
	   		fmt.Printf("v.Name(): %v\n", v.Name())

	   	} */

	//定位
	f, _ := os.Open("a.txt")
	f.Seek(3, 0) //相对索引0（即开头）偏移3位（即从索引3开始读）
	/* Seek设置下一次读/写的位置.offset为相对偏移量，而whence决定
	相对位置：0为相对文件开头，1为相对当前位置，2为相对文件末尾.
	它返回新的偏移量（相对开头）和可能的错误 */
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()
}

func main() {
	// OpenClose()
	// createFile()
	readOps()

}
