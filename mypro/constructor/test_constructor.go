/*
	golang构造函数

golang没有构造函数的概念,可以使用函数模拟构造函数的功能
*/
package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	email string
}

func NewPerson(name string, age int, email string) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be less than 0")
	}
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	return &Person{name: name, age: age, email: email}, nil
}
func main() {
	per, err := NewPerson("klaus", -1, "pilw818@gmail.com")
	if err == nil {
		fmt.Printf("per:%v\n", per)
	} else {
		fmt.Printf("err:%v\n", err)
	}

}
