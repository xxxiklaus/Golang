package main

import "fmt"

func test1() {
	m1 := map[string]string{"name": "klaus", "age": "22", "email": "pilw818@gmail.com"}
	/* 	for k := range m1 {
		fmt.Printf("k: %v\n", k)

	} */
	/* 	for k, v := range m1 {
		fmt.Printf("%v:%v\n", k, v)
	} */
	for k, v := range m1 {
		fmt.Printf("%v : %v\n", k, v) // forr + tab

	}
}
func main() {
	test1()
}
