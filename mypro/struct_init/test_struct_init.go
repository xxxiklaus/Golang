/*
	golang结构体的初始化

未初始化的结构体，成员都是零值 int 0 float 0.0 bool false string nil nil
*/
package main

import "fmt"

func main() {
	type Person struct {
		id, age     int
		name, email string
	}
	var klaus Person
	klaus = Person{
		id:    101,
		age:   20,
		name:  "klaus",
		email: "pilw818@gmail.com",
	}
	fmt.Printf("klaus: %v\n", klaus)

}
