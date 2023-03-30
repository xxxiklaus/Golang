/*
	golang标准库os模块-File文件写操作

下面介绍File结构体相关的文件写操作
*/
package main

import "os"

func write() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755) //O_TRUNC 覆盖
	f.Write([]byte("hello golang"))
	f.Close()
}

func writeString() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755) // O_APPEND 追加
	f.WriteString("hello xxxiklaus")
	f.Close()
}

func writeAt() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("aaa"), 3)
	f.Close()

}

func main() {
	// write()
	// writeString()
	writeAt()

}
