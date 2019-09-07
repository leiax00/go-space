package main

import (
	"fmt"
)

type student interface {
	getName() string
}

type visitorInter interface {
	accept(vi *visitor) string
}

type linTao struct {
	name string
}

var _ visitorInter = new(linTao)

func (linTao *linTao) getName() string {
	fmt.Println("linTao getName")
	return linTao.name
}

func (linTao *linTao) accept(vi *visitor) string {
	fmt.Println("linTao getAge")
	return vi.visit(linTao)
}

type hanMeiMei struct {
	name string
}

var _ visitorInter = new(hanMeiMei)    //强制要求hanMeiMei实现接口visitorInter, 编译时检查
var _ visitorInter = (*hanMeiMei)(nil) //强制要求hanMeiMei实现接口visitorInter，运行时检查

func (hanMeiMei *hanMeiMei) getName() string {
	fmt.Println("hanMeiMei getName")
	return hanMeiMei.name
}

func (hanMeiMei *hanMeiMei) accept(vi *visitor) string {
	fmt.Println("hanMeiMei getAge")
	return vi.visit(hanMeiMei)
}

type liLei struct {
	name string
}

//var _ visitorInter = new(liLei)  // ide报错或编译报错，未实现visitorInter接口

func (liLei *liLei) getName() string {
	fmt.Println("liLei getName")
	return liLei.name
}

type visitor struct {
}

func (vi *visitor) visit(st student) string {
	return st.getName()
}

func main() {
	tao := &linTao{"a"}
	var list []student
	list = append(list, tao, &hanMeiMei{"q"}, &linTao{"b"}, &hanMeiMei{"w"}, &liLei{"z"})
	for _, item := range list {
		fmt.Println(item.getName())
	}

	for _, item := range list {
		if tmp, ok := item.(visitorInter); ok {
			fmt.Println(tmp.accept(&visitor{}))
		}
	}
}
