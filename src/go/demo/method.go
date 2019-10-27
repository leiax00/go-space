package demo

import (
	"fmt"
	"strings"
)

type inter interface {
	say() string
	setName(name string)
}

type ChineseP struct {
	Fun
	Name string
	Age  int
}

func (c *ChineseP) setName(name string) {
	c.Name = name
}

func (c *ChineseP) say() string {
	return "你好，世界！！！"
}

type UsaP struct {
	Fun
	Name string
}

func (u *UsaP) setName(name string) {
	u.Name = name
}

func Main(str string, v []inter) {
	v[0].setName(str)
	fmt.Println(str, v[0].say())
}

type Fun func(...string) string

func (f *Fun) say() string {
	return "Hello, world!!!"
}

func FunC1(strs ...string) string {
	return strings.Join(strs, "~")
}

func Main1(str string, f Fun) {
	fmt.Println(f(str, "you are smart!!!!"))
	fmt.Println(f.say())
}

func aaa() {

}

func Demo() {
	var arr = []inter{&ChineseP{}}
	var arr2 = []inter{&UsaP{}}
	Main("Oh", arr)
	Main("Oh", arr2)
	Main1("Yee", FunC1)
}
