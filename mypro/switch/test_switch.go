//go语言中的switch语句，可以非常容易的判断多个值的情况
//go语言中switch语法
/* switch var1{
    case val1:
	...
	case val2：
	...
	default：
	...
} */
//go语言中switch语句的注意事项
/* 1.支持多条件匹配
2.不同的case之间不使用 break 分隔，默认只会执行一个case
3.如果想要执行多个case，需要使用fallthrough关键字，也可用break终止
4.分支还可以使用表达式 eg：a>10 */

package main

import "fmt"

func f() {
	grade := 'A'
	switch grade {
	case 'A':
		fmt.Println("优秀")
	case 'B':
		fmt.Println("良好")
	default:
		fmt.Println("其他")
	}
}

// 多条件匹配
// go语言switch语句，可以同时匹配多个条件，中间用逗号分隔，其中有一个匹配成功即可
func f1() {
	day := 78
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Println("非法输入...")
	}
}

// case可以是条件表达式（逻辑表达式）
func f2() {
	score := 90
	switch {
	case score >= 90:
		fmt.Println("享受lol、csgo")
	case score >= 80 && score < 90:
		fmt.Println("lol还是往后稍稍,多学习吧")
	default:
		fmt.Println("只要学不死，就往死里学")
	}
}

//fallthrough可以执行满足条件的下一个case
func f3() {
	num := 100
	switch num {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("nothing")
	}
}
func main() {
	// f()
	// f1()
	// f2()
	f3()

}
