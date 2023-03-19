/*
	golang map

map是一种key:value键值对的数据结构容器。map内部实现是哈希表{hash}
map最重要的一点是通过key来快速检索数据，key类似于索引，指向数据的值
map是引用类型的

map语法格式
可以使用内键函数make也可以使用map关键字来定义map

使用make函数
map_variable = make(map[key_data_type]value_data_type)

声明变量，默认map是nil
var map_variable map[key_data_type]value_data_type
map_variable map变量名称
key_data_type key的数据类型
value_data_type 值的数据类型
*/
package main

import "fmt"

func test1() {
	// 类型的声明
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

//声明保存个人信息

func test2() {
	var m1 = map[string]string{"name": "klaus", "age": "22", "email": "pilw818@gmail.com"} //map
	fmt.Printf("m1: %v\n", m1)                                                             //map是无序的

	m2 := make(map[string]string) //make
	m2["name"] = "klaus"
	m2["age"] = "22"
	m2["email"] = "pilw818@gmail.com"

	fmt.Printf("m2: %v\n", m2)
}

//访问map
func test3() {
	m1 := map[string]string{"name": "klaus", "age": "22", "email": "pilw818@gmail.com"}
	var key = "name"

	var value = m1[key]
	fmt.Printf("value: %v\n", value)

}

//判断某个键是否存在
// value, ok := map[key] 如果ok为true，存在，否之，不存在

func test4() {
	var m1 = map[string]string{"name": "klaus", "age": "22", "email": "pilw818@gmail.com"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("------")
	v, ok = m1[k2]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)

}

func main() {
	// test1()
	// test2()
	// test3()
	test4()

}
