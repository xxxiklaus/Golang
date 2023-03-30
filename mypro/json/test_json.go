/*
	golang标准库encoding/json

这个包可以实现json的编码和解码,就是将json字符串转换为struct,或者将struct转换为json.
核心的两个函数
func Marshal(v interface{}) ([]byte, error) //将struct编码成json,可以接收任意类型

func Unmarshal(data []byte, v interface{}) error // 将json转码成结构体

核心的两个结构体
/从输入流读取并解析json

	type Decoder struct {
		/ contains filtered or unexported fields
	}

/写json到输出流

	type Encoder struct {
		/ contains filtered or unexported fields
	}
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func test() { // 结构体转换为json
	p := Person{
		Name:  "klaus",
		Age:   22,
		Email: "pilw818@gmail.com",
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))
}

func test1() { //json转为结构体
	b := []byte(`{"Name":"klaus","Age":22,"Email":"pilw818@gmail.com"}`) //注:此处用反引号 代表多个字符串
	var p Person
	json.Unmarshal(b, &p)

	fmt.Printf("p: %v\n", p)

}

func test2() { //解析嵌套类型
	b := []byte(`{"Name":"klaus","Age":22,"Email":"pilw818@gmail.com","Parents":["big klaus","aster"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("f: %v\n", f)

	for k, v := range f.(map[string]interface{}) {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)

	}

}

func test3() { //解析嵌套引用类型
	type Person struct {
		Name    string
		Age     int
		Email   string
		Parents []string
	}

	p := Person{
		Name:    "klaus",
		Age:     22,
		Email:   "pilw818@gmail.com",
		Parents: []string{"big klaus", "aster"},
	}

	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b))

}

// io流Reader Writer 可以扩展到http websocket 等场景
func test4() {
	f, _ := os.Open("a.json")
	defer f.Close()
	dec := json.NewDecoder(f)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return

		}

		fmt.Printf("v: %v\n", v)
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}

	}

}

func test7() {
	type Person struct {
		Name    string
		Age     int
		Email   string
		Parents []string
	}

	p := Person{
		Name:    "klaus",
		Age:     22,
		Email:   "pilw818@gmail.com",
		Parents: []string{"big klaus", "aster"},
	}

	f, _ := os.OpenFile("a.json", os.O_WRONLY, 0777)
	defer f.Close()

	e := json.NewEncoder(f)
	e.Encode(p)

}

func main() {
	// test()
	// test1()
	// test2()
	// test3()
	test4()

}
