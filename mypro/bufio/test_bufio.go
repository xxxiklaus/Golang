/*
	golang标准库bufio

bufio包实现了有缓冲的I/O.它包装一个io.Reader或者io.Writer接口对象
创建另一个也实现了该接口,且同时还提供了缓冲和一些文本I/O的帮助函数的对象
常量
const (

	defaultBufSize = 4096

)
变量
var (

	ErrInvalidUnreadByte = errors.New("bufio: invaild use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invaild use of UnreadRune")
	ErrBufferFull = errors.New("bufio: buffer full")
	ErrNegativeCount = errors.New("bufio: negative count")

)

var (

	ErrTooLong = errors.New("bufio.Scanner: token too long")
	ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
	ErrAdvanceTooFar = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")

)

	type Reader struct {
		buf          []byte
		rd           io.Reader // reader provided by the client
		r, w         int       // buf read and write positions
		err          error
		lastByte     int // last byte read for UnreadByte; -1 means invalid
		lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
	}  //Reader实现了给一个io.Reader接口对象附加缓冲

func NewReader
func NewReader(rd io.Reader) *Reader
NewReader 创建一个具有默认大小缓冲、从r读取*Reader.NewReader相当于NewReaderSize(rd,4096)

func NewReaderSize
func NewReaderSize(rd io.Reader, size int) *Reader
NewReaderSize创建一个具有最少有size尺寸的缓冲区、从r中读取的Reader
如果参数r已经是一个具有足够大缓冲的Reader类型值，会返回r.

func(*Reader) Rest(r io.Reader)
func (b *Reader) Reset(r io.Reader)
Rest丢弃缓冲中的数据,清除任何错误,将b重设为其他下层从r读取数据.
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func testNewReader() {
	r := strings.NewReader("hello Gopher!")
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)

}

// func(*Reader)Read
func test1() { //func(b*Reader)Read(p[]byte)(n int,err error)
	//Read读取数据写入p.本方法返回写入p的字节数.本方法一次调用最多会调用
	//下层Reader接口一次Read方法,因此返回值n可能小于len(p).读取到达结尾
	//时,返回值n将为0而err将为io.EOF
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	for {
		n, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p): %v\n", string(p[0:n])) //读了几个
		}
	}
}

// func(*Reader)ReadByte  func (b*Reader) ReadByte() (c byte,err error)
// ReadByte 读取并返回一个字节.如果没有可用的数据,会返回错误

// func(*Reader)UnreadByte  func (b*Reader) UnreadByte() error
// UnreadByte吐出最近一次读取操作读取的最后一个字节.(只能吐出最后一个,多次调用会出问题)
func test2() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte()
	fmt.Printf(" %c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf(" %c\n", c)

	br.UnreadByte() //读B
	c, _ = br.ReadByte()
	fmt.Printf(" %c\n", c)

}

// func (*Reader) ReadRune  func (b *Reader) ReadRune() (r rune, size int, err error)
// ReadRune 读取一个utf-8编码的unicode码值,返回该码值、其编码长度和可能的错误
// 如果utf-8编码非法,读取位置只移动一字节,返回U+FFFD,返回值size为1而err为nil
// 如果没有可用的数据,会返回错误

// func (*Reader) UnreadRune func (b *Reader) UnreadRune() error
// UnreadRunc吐出最近一次ReadRune调用读取的unicode编码的值.如果最近一次读取不是
// 调用的ReadRune,会返回错误.(由此看出,UnreadRune比UnreadByte严格很多)
func test3() {
	s := strings.NewReader("你好,世界！")
	br := bufio.NewReader(s)

	c, size, _ := br.ReadRune()
	fmt.Printf(" %c %v\n", c, size)

	c, size, _ = br.ReadRune()
	fmt.Printf(" %c %v\n", c, size)

	br.UnreadRune()
	c, size, _ = br.ReadRune()
	fmt.Printf(" %c %v\n", c, size)
}

// func (*Reader) ReadLine func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
//ReadLine是一个低水平的行数据读取原语.大多数调用者应使用ReadBytes('\n') 或 ReadString('\n')代替
// 或使用 Scanner.

//ReadLine尝试返回单行数据,不包括行尾标志的字节.如果该行太长超过了缓冲区,返回值isPrefix会被设为true
// 并返回行的前面一部分.该行剩下的部分将在之后的调用中返回.返回值isPrefix会在返回该行最后一个片段时
//才设为false.返回切片是缓冲的子切片,仅在下一次读取操作之前有效.ReadLine要么返回一个非nil的line
//要么返回一个非nil的err两个返回值至少一个非nil

//从ReadLine返回的文本不包含行尾的标志字节(“\r\n”或“\n”)。如果输入流结束时没有行尾标志字节
// 方法不会出错也不会指出这一情况.在调用ReadLine之后调用UnreadByte会总是吐出最后一个读取的字节
// (很可能是该行的行尾的标志字节)即使该字节不是ReadLine返回值的一部分

func test4() {
	s := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf(" %q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf(" %q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf(" %q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf(" %q %v\n", w, isPrefix)

}

// func (*Reader) ReadSlice  func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

// ReadSlice读取直到第一次遇到delim字节,返回缓冲里的包含已读取的数据和delim字节的切片.
// 该返回值只在下一次读取操作之前合法.如果ReadSlice放在在读取到delim之前遇到错误,它会返回
// 在错误之前读取的数据在缓冲中的切片以及该错误(通常是 io.EOF).如果在读取到delim之前缓冲
// 就被写满了，ReadSlice将失败并返回ErrBufferFull.因为ReadSlice的返回值将被下一次I/O操作重写
// 调用者应尽量使用ReadBytes或ReadString替代本法功法.当且仅当ReadBytes方法返回的切片不以delim
// 结尾时,ReadSlice会返回一个非nil的错误
func test5() {
	s := strings.NewReader("ABC,DEF,GHI,JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadSlice(',')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf(" %q\n", w)
}

// func (*Reader) ReadBytes func (b *Reader) ReadBytes(delim byte) ([]byte, error)
// ReadBytes读取直到第一次遇到delim字节,返回一个包含已读取的数据和delim字节的切片.
// 如果ReadBytes放在在读取到delim之前遇到错误,它会返回在错误之前读取的数据在缓冲中
// 的切片以及该错误(通常是 io.EOF).当且仅当ReadBytes方法返回的切片不以delim结尾时
// ReadBytes会返回一个非nil的错误

func test6() {
	s := strings.NewReader("ABC,DEF,GHI,JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadBytes(',')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadBytes(',')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadBytes(',')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadBytes(',')
	fmt.Printf(" %q\n", w)
}

// func (*Reader) ReadString func (b *Reader) ReadString(delim byte) (string, error)
// ReadString读取直到第一次遇到delim字节,返回一个包含已读取的数据和delim字节的字符串.
// 如果ReadString放在在读取到delim之前遇到错误,它会返回在错误之前读取的数据在缓冲中
// 的切片以及该错误(通常是 io.EOF).当且仅当ReadString方法返回的切片不以delim结尾时
// ReadString会返回一个非nil的错误

func test7() {
	s := strings.NewReader("ABC DEF GHI JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadString(' ')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf(" %q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf(" %q\n", w)
}

// func (*Reader) WriteTo  func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
// WriteTo方法实现了io.WriterTo接口

func test8() {
	s := strings.NewReader("ABCDEFGHIJKLMN")
	br := bufio.NewReader(s)
	b := bytes.NewBuffer(make([]byte, 0))

	br.WriteTo(b)
	fmt.Printf("b: %s\n", b)
}

// func NewWriter func NewWriter(w io.Writer) *Writer
// NewWriter returns a new Writer whose buffer has the default size.
// If the argument io.Writer is already a Writer with large enough
// buffersize, it returns the underlying Writer.

func test9() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777) // File Reader Writer 接口均被实现

	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("hello,Gophers!")
	w.Flush()
}

// func NewWriterSize  func NewWriterSize(w io.Writer, size int) *Writer
// NewWriterSize returns a new Writer whose buffer has at least the specified
// size.If the argument io.Writer is already a Writer with large enough size,
// it returns the underlying Writer.

// func (*Writer) Reset func (b *Writer) Reset(w io.Writer)
// Reset discards any unflushed buffered data, clears any error, and resets
// b to write its output to w. Calling Reset on the zero value of Writer
// initializes the internal buffer to the default size.

func test10() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123456789")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)
	fmt.Println(c)
}

// func (*Writer) Buffered func (b *Writer) Buffered() int
// Buffered returns the number of bytes that have been written into the current buffer.

// func (*Writer) Available func (b *Writer) Available() int
// Available returns how many bytes are unused in the buffer.

// func (*Writer) AvailableBuffer func (b *Writer) AvailableBuffer() []byte
// AvailableBuffer returns an empty buffer with b.Available() capacity. This
// buffer is intended to be appended to and passed to an immediately succeeding
// Write call. The buffer is only valid until the next write operation on b.

func test11() {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	w.Flush()
}

// func (*Writer) Write func (b *Writer) Write(p []byte) (nn int, err error)
// Write writes the contents of p into the buffer. It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining why the write is short.

// func (*Writer) WriteByte func (b *Writer) WriteByte(c byte) error
// WriteByte writes a single byte.

// func (*Writer) WriteRune  func (b *Writer) WriteRune(r rune) (size int, err error)
// WriteRune writes a single Unicode code point, returning the number of bytes written and any error.

func test12() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	//写入缓存
	//byte等同于 int8
	bw.WriteByte('H')
	bw.WriteByte('e')
	bw.WriteByte('l')
	bw.WriteByte('l')
	bw.WriteByte('o')
	bw.WriteByte(' ')
	//rune等同于 int32
	bw.WriteRune('世')
	bw.WriteRune('界')
	bw.WriteRune('!')
	//写入b
	bw.Flush()
	fmt.Println(b)

}

// func (*Writer) WriteString  func (b *Writer) WriteString(s string) (int, error)
// WriteString writes a string. It returns the number of bytes written. If the count
// is ess than len(s), it also returns an error explaining why the write is short.

// func (*Writer) ReadFrom  func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
// ReadFrom implements io.ReaderFrom. If the underlying writer supports the ReadFrom
// method,this calls the underlying ReadFrom. If there is buffered data and an
// underlying ReadFrom, this fills the buffer and writes it before calling ReadFrom.

func test13() {
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("Hello 世界!")
	bw := bufio.NewWriter(b)
	bw.ReadFrom(s)
	// bw.Flush() ReadFrom 无需使用Flush,其自己已经写入.
	fmt.Println(b)
}

// func (*Writer) Flush  func (b *Writer) Flush() error
// Flush writes any buffered data to the underlying io.Writer.

func test14() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) //4096
	fmt.Println(bw.Buffered())  //0

	bw.WriteString("ABCDEFGHIJKLMN")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("b: %q\n", b)

	bw.Flush() //刷新
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("b: %q\n", b)

}

/* type ReadWriter
type ReadWriter struct {
	/ *Reader
	/ *Writer
} */
//ReadWriter类型保管了指向Reader和Writer类型的指针,实现了io.ReaderWriter接口

// func NewReadWriter func NewReadWriter(r *Reader, w *Writer) *ReadWriter
//NewReadWriter 申请创建一个新的\将读写操作分派给r和w 的ReadWriter

func test15() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')
	fmt.Println(string(p)) //123
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b)
}

// func (*Scanner) Split  func (s *Scanner) Split(split SplitFunc)
// Split sets the split function for the Scanner. The default split function is ScanLines.
// Split panics if it is called after scanning has started.

func test16() {
	s := strings.NewReader("ABC DEF GHI JKL")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords) //以空格作为分隔符进行分割
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

// func (*Scanner) Scan func (s *Scanner) Scan() bool
// Scan advances the Scanner to the next token, which will then be available through
// the Bytes or Text method. It returns false when the scan stops, either by reaching
// the end of the input or an error. After Scan returns false, the Err method will return
// any error that occurred during scanning, except that if it was io.EOF,Err will return nil.
// Scan panics if the split function returns too many empty tokens without advancing the
// input. This is a common error mode for scanners.

func test17() {
	s := strings.NewReader("Hello 世界!")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	}
}

// func (*Scanner) Bytes func (s *Scanner) Bytes() []byte
// Bytes returns the most recent token generated by a call to Scan. The underlying
// array may point to data that will be overwritten by a subsequent call to Scan.
// It does no allocation.

// func (*Scanner) Text func (s *Scanner) Text() string
// Text returns the most recent token generated by a call
// to  Scan as a newly allocated string holding its bytes.

// func (*Scanner) Err func (s *Scanner) Err() error
// Err returns the first non-EOF error that was encountered by the Scanner.

func main() {
	// testNewReader()

	// test1()
	// test2()

	// test3()
	// test4()

	// test5()
	// test6()

	// test7()
	// test8()

	// test9()
	// test10()

	// test11()
	// test12()

	// test13()
	// test14()
	// test15()

	// test16()
	test17()

}
