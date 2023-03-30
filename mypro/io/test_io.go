/*
	golang标准库io包 input output

Go语言中,为了方便开发者使用,将IO操作封装在了如下几个包中:
io为IO原语(I/O primitives)提供基本的接口 os File Reader Writer
io/ioutil封装一些实用的I/O函数
fmt实现格式化I/O,类似于C语言中的printf 和 scanf format
bufio实现带缓冲I/O函数
io —— 基本的IO接口
在io包中最重要的是两个接口: Reader 和 Writer 接口.
本章所提到的各种IO包都跟这两个接口有关,即实现了这两个接口,它就有了IO的功能
Reader接口

	type Reader interface {
		Read(p []byte) (n int,err error)
	}

Writer接口

	type Writer interface {
		Write(p []byte) (n int,err error)
	}
*/

/* 哪些类型实现了Reader和Writer接口
os.File 同时实现了 io.Reader 和 io.Writer
strings.Reader 实现了io.Reader
bufio.Reader/Writer  分别实现了 io.Reader / io.Writer
bytes.Buffer io.Reader 和 io.Writer
bytes.Reader io.Reader
compress/gzip.Reader/Writer io.Reader / io.Writer
crypto/cipher.StreamReader/StreamWriter io.Reader / io.Writer
crypto/tls.Conn io.Reader 和 io.Writer
encoding/csv.Reader/Writer io.Reader / io.Writer */

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	//file1 file2
	r := strings.NewReader("hello golang")

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func testCopyBuffer() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func testLimitReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

func testMultiReader() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

func testMultiWriter() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 strings.Builder
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}

func testSectionReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
}

func main() {
	/* 	r := strings.NewReader("hello golang")
	   	buf := make([]byte, 20)
	   	r.Read(buf)
	   	fmt.Printf("string(buf): %v\n", string(buf)) */

	// testCopy()
	// testCopyBuffer()
	// testLimitReader()
	// testMultiReader()
	// testMultiWriter()
	testSectionReader()

}
