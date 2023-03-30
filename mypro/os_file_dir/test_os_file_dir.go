/*
	golang标准库os模块-文件目录相关

os标准库实现了平台(操作系统)无关的编程接口
https://pkg.go.dev/std
*/
package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func makeDir() {
	/* err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */
	err1 := os.MkdirAll("a/b/c", os.ModePerm)
	if err1 != nil {
		fmt.Printf("err1: %v\n", err1)
	}
}

// 删除目录或者文件
func remove() {
	/* err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */
	err := os.RemoveAll("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 获得工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err : %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}

}

// 修改工作目录
func chWd() {
	err := os.Chdir("d:/")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}

// 获得临时目录
func getTemp() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// 重命名文件
func renameFile() {
	err := os.Rename("test.txt", "test1.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func readFile() {
	b, err := os.ReadFile("test1.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:])) //取字节的切片
	}
}

// 写文件
func writeFile() {
	s := "hello golang"
	os.WriteFile("test1.txt", []byte(s), os.ModePerm)
}
func main() {
	// createFile()
	makeDir()
	// remove()
	// getWd()
	// chWd()
	// getTemp()
	// renameFile()
	// readFile()
	// writeFile()

}
