/*
	golang标准库ioutil包

ReadAll 读取数据,返回读到的字节 slice
ReadDir 读取一个目录,返回目录入口数组 []os.FileInfo
ReadFile 读取一个文件,返回文件内容(字节slice)
WriteFile 根据文件路径,写入字节slice
TempDir 在一个目录中创建指定前缀名的临时目录,返回新临时目录的路径
TempFile 在一个目录中创建指定前缀的临时文件,返回os.File
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func testReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}

func testReadDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}

func testReadFile() {
	b, _ := ioutil.ReadFile("a.txt")
	fmt.Println(string(b))
}

func testWriteFile() {
	message := []byte("hello,Gophers!")
	err := ioutil.WriteFile("a.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func testTempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tmpfile.Name(): %v\n", tmpfile.Name())

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// testReadAll()
	// testReadDir()
	// testReadFile()
	// testWriteFile()
	testTempFile()

}
