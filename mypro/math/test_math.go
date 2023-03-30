/*
	golang标准库math

该包包含一些常量和一些有用的数学计算函数,eg:三角函数、随机数、绝对值、平方根等
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func test() { //常量
	fmt.Printf("float64的最大值是:%.f\n", math.MaxFloat64)
	fmt.Printf("float64的最小值是:%.f\n", math.SmallestNonzeroFloat64)
	fmt.Printf("float32的最大值是:%.f\n", math.MaxFloat32)
	fmt.Printf("float32的最小值是:%.f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("math.MaxInt8: %d\n", math.MaxInt8)
	fmt.Printf("math.MinInt8: %d\n", math.MinInt8)
	fmt.Printf("math.MaxUint8: %d\n", math.MaxUint8)
	fmt.Printf("math.MaxInt16: %d\n", math.MaxInt16)
	fmt.Printf("math.MinInt16: %d\n", math.MinInt16)
	fmt.Printf("math.MaxUint16: %d\n", math.MaxUint16)
	fmt.Printf("math.MaxInt32: %d\n", math.MaxInt32)
	fmt.Printf("math.MinInt32: %d\n", math.MinInt32)
	fmt.Printf("math.MaxUint32: %d\n", math.MaxUint32)
	fmt.Printf("math.MaxInt64: %d\n", math.MaxInt64)
	fmt.Printf("math.MinInt64: %d\n", math.MinInt64)
	fmt.Printf("math.Pi: %.200f\n", math.Pi)
}

func test1() { //绝对值 Abs(x float64) float64
	fmt.Printf("math.Abs(-3.14): [%.2f]\n", math.Abs(-3.14))
	// }

	// func Pow(x, y float64) float64 { //取x的y次方
	fmt.Printf("[2]的16次方为: [%.f]\n", math.Pow(2, 16))
	// }

	// func Pow10(n int) float64 { //取立方
	fmt.Printf("math.Pow10(3): [%.f]\n", math.Pow10(3)) //10的[3]次方为
	// }

	// func Sqrt(x float64) float64 { //取x的开平方
	fmt.Printf("math.Sqrt(64): [%.f]\n", math.Sqrt(64))
	// }

	// func Cbrt(x float64) float64 { //取x的开立方
	fmt.Printf("math.Cbrt(27): [%.f]\n", math.Cbrt(27))
	// }

	// func Ceil(x float64) float64 { //向上取整
	fmt.Printf("math.Ceil(3.14): [%.f]\n", math.Ceil(3.14))
	// }

	// func Floor(x float64) float64 { //向下取整
	fmt.Printf("math.Floor(8.75): [%.f]\n", math.Floor(8.75))
	// }

	// func Mod(x, y float64) float64 { //取余数
	fmt.Printf("[10/3]的余数为: [%.f]\n", math.Mod(10, 3))
	// }

	// func Modf(f float64) (int float64, frac float64) { //分别取整数部分 小数部分
	Integer, Decimal := math.Modf(3.14159265358979)
	fmt.Printf("Integer: [%.f]", "Decimal:[%.14f]\n", Integer, Decimal)
	// }
}

func init() { //随机数
	//seed 种子 以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
	// rand.Seed(1)
}

func main() {
	// test()
	// test1()

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println(a)
	}

	fmt.Println("-------")
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}

	fmt.Println("-------")
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
	// fmt.Printf("rand.Int: %d\n", rand.Int)

}
