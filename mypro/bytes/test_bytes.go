/*
	golang标准库bytes

bytes包提供了对字节切片进行读写操作的一系列函数,字节切片处理的函数比较多分为
基本处理函数,比较函数,后缀检查函数,索引函数,大小写处理函数和子切片处理函数等
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func testTrans() { //强制转换
	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)

	var s string = "hello golang"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)

}

func testContains() { //是否包含
	s := "xxxiklaus818.com"
	b := []byte(s)
	b1 := []byte("xxxiklaus")
	b2 := []byte("Xxxiklaus")

	s1 := strings.Contains("hello golang", "hello")
	fmt.Printf("s1: %v\n", s1)

	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)

	b3 = bytes.Contains(b, b2)
	fmt.Printf("b3: %v\n", b3)

}

func testCount() { //计数
	s := []byte("Chooooooooooovy!")
	sep1 := []byte("C")
	sep2 := []byte("h")
	sep3 := []byte("o")

	fmt.Println(bytes.Count(s, sep1))
	fmt.Println(bytes.Count(s, sep2))
	fmt.Println(bytes.Count(s, sep3))

}

func testRepeat() { //重复
	b := []byte("hi")
	fmt.Println(string(bytes.Repeat(b, 1)))
	fmt.Println(string(bytes.Repeat(b, 3)))

}

func testReplace() { //替换
	s := []byte("hello golang")
	old := []byte("o")
	news := []byte("xx")
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  //不替换
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  //1次
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  //2次
	fmt.Println(string(bytes.Replace(s, old, news, -1))) //替换若干次

}

func testRunes() {
	s := []byte("你好世界") //一个汉字表示3个字节
	r := bytes.Runes(s)
	fmt.Println("转换前字符串的长度:", len(s))
	fmt.Println("转换后字符串的长度:", len(r))

}

func testJoin() { //连接
	s1 := [][]byte{[]byte("你好"), []byte("世界")}
	sep1 := []byte(" ")
	fmt.Println(string(bytes.Join(s1, sep1)))
	sep2 := []byte(",")
	fmt.Println(string(bytes.Join(s1, sep2)))
	sep3 := []byte("#")
	fmt.Println(string(bytes.Join(s1, sep3)))

}

//Reader类型
//Reader实现了io.Reader,io.ReadAt,io.WriterTo,io.Seeker.io.ByteScanner
// io.RuneScanner等接口,Reader是只读的,可以seek

func testReader() {
	data := "123456789"
	re := bytes.NewReader([]byte(data)) //通过[]byte创建Reader
	fmt.Println("re len :", re.Len())   //返回未读取部分的长度
	fmt.Println("re size :", re.Size()) //返回底层数据总长度
	fmt.Println("-------")

	buf := make([]byte, 2)
	for {
		n, err := re.Read(buf) //读取数据
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
	fmt.Println("-------")
	re.Seek(0, 0) //设置偏移量,因为上面的操作已经修改了读取位置等信息
	for {
		b, err := re.ReadByte() //一个字节一个字节的读
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}
	fmt.Println("--------")

	re.Seek(0, 0)
	off := int64(0)
	for {
		n, err := re.ReadAt(buf, off) //指定偏移量读取
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}

// Buffer类型
// 缓冲区是具有读取和写入方法的可变大小的字节缓冲区.Buffer的零值是准备使用的空缓冲区.

/* 声明Buffer的方法 */
// var b bytes.Buffer // 直接定义一个Buffer变量,不用初始化,可直接使用
// b := new(bytes.Buffer) //使用New返回Buffer变量
// b := bytes.NewBuffer(s []byte) //从一个[]byte切片,构造一个Buffer
// b := bytes.NewBufferString(s string) //从一个string变量,构造一个Buffer
func testBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1: %v\n", b1)
	var b2 = bytes.NewBuffer([]byte("hello"))
	fmt.Printf("b2: %v\n", b2)

}

/* 往Buffer中写入数据 */
// b.Write(d []byte) //将切片d 写入Buffer的尾部
// b.WriteString(s string) //将字符串s写入Buffer尾部
// b.WriteByte(c byte) //将字符c写入Buffer尾部
// b.WriteRune(r rune) //将一个rune类型的数据放到缓冲区的尾部
// b.WriteTo(w io.Writer) //将Buffer中的内容输出到实现了io.Writer接口的可写入对象中
func testBuffer1() {
	var b bytes.Buffer
	n, _ := b.WriteString("hello")
	fmt.Printf("n:  %v\n", n)
	fmt.Printf("b: %v\n", string(b.Bytes()))

}

/* 从Buffer中读取数据到指定容器 */
// c := make([]byte,8)
// b.Read(c) //一次读取8个byte到c容器中,每次读取新的8个byte覆盖c中原来的内容
// b.ReadByte() // 读取第一个byte,b的第一个byte被拿掉,赋值给 a => a,_ := b.ReadByte()
// b.ReadRune() // 读取第一个rune,b的第一个rune被拿掉,赋值给 r => r,_ := b.ReadRune()
// b.ReadBytes(delimiter byte) / 需要一个byte作为分隔符 , 读的时候从缓冲器里找第一个出现的分隔符
//找到后,把从缓冲器头部开始到分隔符之间的所有byte进行返回,作为byte类型的slice,返回后,缓冲器也会空掉一部分
// b.ReadString(delimiter byte) /需要一个byte作为分隔符 , 读的时候从缓冲器里找第一个出现的分隔符
//找到后,把从缓冲器头部开始到分隔符之间的所有byte进行返回,作为字符串返回,返回后,缓冲器也会空掉一部分
// b.ReadFrom(i io.Reader) /从一个实现io.Reader接口的r,把r里的内容读到缓冲器里 , n 返回读的数量

/* 注:若是将文件中的内容写入Buffer,则使用ReadFrom(io.Reader) */
func testBuffer2() {
	var b = bytes.NewBufferString("hello golang")
	b1 := make([]byte, 2)

	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("b1: %v\n", string(b1[0:n]))
	}

}

func main() {
	// testTrans()
	// testContains()
	// testCount()
	// testRepeat()
	// testReplace()
	// testRunes()
	// testJoin()
	// testReader()
	// testBuffer()
	// testBuffer1()
	testBuffer2()

}
