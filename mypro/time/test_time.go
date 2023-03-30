/* golang标准库time
time包提供测量和显示时间的功能. */
// 基本使用 打印显示出现在的时间
package main

import (
	"fmt"
	"time"
)

func nowtime() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	//获取当前时间 current time: 2023-03-29 22:02:17.1259167 +0800 CST m=+0.013017501
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d-%02d-%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T,%T\n", now, year, month, day, hour, minute, second)
	//time.Time,int,time.Month,int,int,int,int

}

//时间戳
// 在编程中对于时间戳的应用尤为广泛,例如在web开发中做cookies有效期,接口加密
// Redis中的key有效期等等,大部分都是使用到了时间戳
// 时间戳是自1970年1月1日(08:00:00GMT)至当前时间的总毫秒数.
// 它也被称为Unix时间戳(UnixTimestamp)
/* Golang中获取时间戳 */

func timeStamp() {
	now := time.Now()
	fmt.Printf("TimeStamp type: %T, TimeStamp: %v\n", now.Unix(), now.Unix())
	//当前时间戳 TimeStamp type: int64, TimeStamp: 1680099107
	fmt.Printf("TimeStamp type: %T, TimeStamp: %v\n", now.UnixNano(), now.UnixNano()) //纳秒时间戳
}

//时间戳转化为普通时间的时间格式 go语言中用time.Unix 实现瞬间替换

func timeStampTOTime() {
	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d-%02d-%02d:%02d\n", year, month, day, hour, minute, second)

}

/* 操作时间 */
func timeAdd() { //Add
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)
}

func timeSub() { //Sub
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", difference) //目标时间与此时相比差
}

// func (t Time) Before(u Time) bool / 如果t代表的时间点在u之前,返回true,否则返回false
// func (t Time) Equal(u Time) bool /判断两个时间是否相同,会考虑时区影响,因此不同时区标准的时间也可以正确比较
// func (t Time) After(u Time) bool /如果t代表的时间点在u之后,返回true,否则返回false

//定时器 使用time.Tick(时间间隔)来设置,本质上是一个通道(channel)
/* func tick() {
	ticker := time.NewTicker(time.Second) //定义一个1秒间隔的定时器
	for i:= range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
} */

// 时间格式化 Format func (t Time) Format(layout string) string
func timeFormat() {
	// Parse a time value from a string in the standard Unix format.
	t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	// time.Time's Stringer method is useful without any format.
	fmt.Println("default format:", t)

	// Predefined constants in the package implement common layouts.
	fmt.Println("Unix format:", t.Format(time.UnixDate))

	// The time zone attached to the time value affects its output.
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))

	fmt.Println("in Shanghai with seconds:", t.In(tz).Format("2006-01-02T15:04:05 -070000"))

	fmt.Println("in Shanghai with colon seconds:", t.In(tz).Format("2006-01-02T15:04:05 -07:00:00"))

	// The rest of this function demonstrates the properties of the
	// layout string used in the format.

	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	// There are some wrinkles illustrated below.

	// Most uses of Format and Parse use constant layout strings such as
	// the ones defined in this package, but the interface is flexible,
	// as these examples show.

	// Define a helper function to make the examples' output look nice.
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	// Print a header in our output.
	fmt.Printf("\nFormats:\n\n")

	// Simple starter examples.
	do("Basic full date", "Mon Jan 2 15:04:05 MST 2006", "Wed Feb 25 11:06:39 PST 2015")
	do("Basic short date", "2006/01/02", "2015/02/25")

	// The hour of the reference time is 15, or 3PM. The layout can express
	// it either way, and since our value is the morning we should see it as
	// an AM time. We show both in one format string. Lower case too.
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	// When parsing, if the seconds value is followed by a decimal point
	// and some digits, that is taken as a fraction of a second even if
	// the layout string does not represent the fractional second.
	// Here we add a fractional second to our time value used above.
	t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}
	// It does not appear in the output if the layout string does not contain
	// a representation of the fractional second.
	do("No fraction", time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")

	// Fractional seconds can be printed by adding a run of 0s or 9s after
	// a decimal point in the seconds value in the layout string.
	// If the layout digits are 0s, the fractional second is of the specified
	// width. Note that the output has a trailing zero.
	do("0s for fraction", "15:04:05.00000", "11:06:39.12340")

	// If the fraction in the layout is 9s, trailing zeros are dropped.
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")
}

// 解析字符串格式的时间
func testTimeString() {
	now := time.Now()
	fmt.Println(now)

	loc, err := time.LoadLocation("Asia/Shanghai") //加载时区
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2023/03/30 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func main() {
	// nowtime()
	// timestamp()
	// timeStampTOTime()
	// timeAdd()
	// timeSub()
	// timeFormat()
	testTimeString()

}
