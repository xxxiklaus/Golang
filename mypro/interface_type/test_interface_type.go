/*
	golang接口和类型的关系

1.一个类型可以实现多个接口
2.多个类型可以实现同一个接口(多态)
*/
package main

import "fmt"

/* 一个类型可以实现多个接口,eg:有一个Player接口
可以播放音乐,有一个Video接口可以播放视频,一个手机Mobile实现这两个
接口,既可以播放音乐又可以播放视频 */

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct {
}

func (m Mobile) playMusic() {
	fmt.Println("DJ drop the beat")
}

func (m Mobile) playVideo() {
	fmt.Println("have some fun video")
}

/* 多个类型实现同一个接口,eg:猫类型Cat和狗类型Dog都可以实现实现该接口
都可以把猫和狗当宠物类型对待,这在其他语言中称作多态
*/

type Pet interface {
	eat()
	sleep()
}

type Cat struct {
}

type Dog struct {
}

func (cat Cat) eat() {
	fmt.Println("cat eat")
}

func (cat Cat) sleep() {
	fmt.Println("cat sleep")
}

func (dog Dog) eat() {
	fmt.Println("dog eat")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep")
}

func main() {
	m := Mobile{}
	m.playMusic()
	m.playVideo()

	/* 	dog := Dog{}
	   	cat := Cat{}
	   	cat.eat()
	   	cat.sleep()
	   	dog.eat()
	   	dog.sleep() */

	var pet Pet
	pet = Dog{}
	pet.eat()
	pet.sleep()
	pet = Cat{}
	pet.eat()
	pet.sleep()

}
