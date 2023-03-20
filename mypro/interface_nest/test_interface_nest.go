/*
	golang接口嵌套

golang中接口可以通过嵌套,创建新的接口.
eg:飞鱼,既可以飞天,又可以水里游.创建一个飞Fly接口
创建一个Swim接口,飞鱼接口有这两个接口组成
*/
package main

import "fmt"

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface { /* 接口的组合 */
	Flyer
	Swimmer
}

type Fish struct {
}

func (fish Fish) fly() {
	fmt.Println("wuhuairlines...")
}

func (fish Fish) swim() {
	fmt.Println("wuhugulugulu...")
}

func main() {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()

}
