//Go语言数字类型
/* Go 语言支持整型和浮点型数字，并且原生支持复数，其中位的运算采用补码。
Go 也有基于架构的类型，例如：int、uint 和 uintptr。
这些类型的长度都是根据运行程序所在的操作系统类型所决定的：
int 和 uint 在 32 位操作系统上，它们均使用 32 位（4 个字节），在 64 位操作系统上，它们均使用 64 位（8 个字节）。
uintptr 的长度被设定为足够存放一个指针即可。
Go 语言中没有 float 类型。（Go语言中只有 float32 和 float64）没有double类型.
与操作系统架构无关的类型都有固定的大小，并在类型的名称中就可以看出来：
*/
/*  整数：
int8(-128到127)
int16(-32768到32767)
int32(-2147483648到2147483647)
int64(-9223372036854775808到9223372036854775807)
无符号整数：
uint8 (0到255)
uint16(0到65535)
uint32(0到4294967295)
uint64(0到18446744073709551615)
浮点型(IEEE-754标准):
float32 (+-le-45->+-3.4*le38)
float64 (+-5*le-324->107*le308) */
// int型 是计算最快的一种类型 (整型的零值为0，浮点型零值为0.0)
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {

	/* 	i := 100
	   	fmt.Printf("%T\n", i) */

	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

	var f32 float32
	var f64 float64

	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(i64), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)

	/* 	var ui uint
	   	ui = uint(math.MaxUint64) // 再+1会导致overflows 错误
	   	fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui) */

	var imax, imin int
	imax = int(math.MaxInt64) // 再+1会导致overflows 错误
	imin = int(math.MinInt64) // 再+1会导致overflows 错误

	fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax)

	// 以二进制、八进制或十六进制浮点数的格式定义数字
	/* 	var a int = 10
	   	fmt.Printf("%d \n", a) //10
	   	fmt.Printf("%b \n", a) //1010 占位符 %b 表示二进制

	   	//八进制 以0开头
	   	var b int = 077
	   	fmt.Printf("%o \n", b) //77

	   	//十六进制 以0x开头
	   	var c int = 0xff
	   	fmt.Printf("%x \n", c) //ff
	   	fmt.Printf("%X \n", c) //FF */
	//浮点型
	/* 	go语言支持两种浮点型数：float32和float64.这两种浮点型数据格式遵循IEEE -754标准const
	   	float32的浮点数最大范围约为3.4e38，可以使用常量定义：math.MaxFloat32.float64的浮点数
	   	最大范围约为1.8e308，可以使用一个常量定义：math.MaxFloat64.
	   	(打印浮点数时，可使用fmt包配合动词%f)  */
	/* fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi) */
	// 复数类型 complex64和complex128

	/* 	var c1 complex64
	   	c1 = 1 + 2i
	   	var c2 complex128
	   	c2 = 2 + 3i

	   	fmt.Println("c1")
	   	fmt.Println("c2")
	*/
}
