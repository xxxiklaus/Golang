/* golang标准库os包和环境相关的方法 */
package main

import (
	"fmt"
	"os"
)

func main() {
	//获得所有环境变量
	s := os.Environ()
	fmt.Printf("s: %v\n", s)
	//打开某个环境变量
	s2 := os.Getenv("GOPATH")
	fmt.Printf("s2: %v\n", s2)
	//设置环境变量
	os.Setenv("env1", "env1")
	s2 = os.Getenv("aaa")
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("-------")

	//查找
	s3, b := os.LookupEnv("env1")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("s3: %v\n", s3)

	//替换
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/user/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	//清空环境变量
	//os.Clearenv()

}
