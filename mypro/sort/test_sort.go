/*
	golang标准库中的sort包

sort包的内容以及使用
sort包提供了排序切片和用户自定义数据集以及相关功能的函数.
sort包主要针对 []int、 []float64、 []string、以及其他自定义切片的排序

结构体 type Float64Slice struct  type IntSlice type StringSlice

	接口 type Interface interface{
		Len() int /Len方法返回集合中的元素个数
		Less(i, j int) bool / i>j,该方法返回索引i的元素是否比索引j的元素小
		Swap(i, j int) / 交换i,j的值
	}
*/
package main

import (
	"fmt"
	"sort"
)

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func testfloat64() {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	fmt.Println(s)
}

func teststring() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)

	for _, v := range s {
		fmt.Println(v, []byte(v))
	}
}

/* type testIntSlice [][]int

func (l testIntSlice) Len() int           { return len(l) }
func (l testIntSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testIntSlice) Less(i, j int) bool { return l[i][1] < l[j][1] } */

/* type testSlice []map[string]float64

func (l testSlice) Len() int           { return len(l) }
func (l testSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool { return l[i]["a"] < l[j]["a"] } //按照"a"对应的值排序
*/

func main() {
	/* 	n := []uint{1, 3, 2}
	   	sort.Sort(NewInts(n))
	   	fmt.Println(n)
	*/
	// testfloat64()
	// teststring()
	/* 	ls := testIntSlice{
		{1, 4},
		{9, 3},
		{7, 5},
	} */
	/* ls := testSlice{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}
	fmt.Println(ls)
	sort.Sort(ls)
	fmt.Println(ls) */
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}
