/*
	golang标准库errors

errors包实现了操作错误的函数.go语言使用error类型来返回函数执行过程中遇到的错误
如果返回的error值为nil,则表示未遇到错误,否则error会返回一个字符串用于说明遇到了什么错误

error结构

	type error interface {
		Error() string
	}

可以用任何类型区实现它(只要添加一个Error方法即可)，也就是说error可以是任何类型
这意味着,函数返回的error值实际可以包含任意信息,不一定是字符串.

error不一定表示一个错误,它可以表示任何信息,比如io包中就用error类型的
io.EOF表示数据读取结束.而不是遇到了什么错误.

error包实现了一个最简单的error类型，只包含一个字符串,它可以记录大多数情况下遇到的
错误信息.errors包的用法也很简单,只有一个New函数,用于生成一个最简单的error对象：
*/
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

// func New(text string) error
func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}

// 自定义错误
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v:%v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(2023, 3, 29, 21, 00, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

// func As(err error, target any) bool
func testAs() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	// f, err := os.Open("a.txt")
	// testAs()
	/* 	s, err := check("hello golang")
	   	if err != nil {
	   		fmt.Printf("err: %v\n", err.Error())
	   	} else {
	   		fmt.Printf("s: %v\n", s)
	   	}*/

	err := oops()
	if err != nil {
		fmt.Println(err)
	}

}
