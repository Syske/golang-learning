package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
	age  int8
}

var u = user{name: "syske", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Recdived Value", pu)
	pu.name = "feng"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

func main() {
	c := make(chan *user, 5)
	c <- g
	fmt.Println(g)
	// modify g
	g = &user{name: "lei", age: 100}
	// 因为 c 是在 g 改变前发送到通道的，所以这里其实不会变
	go printUser(c)
	// 这里是针对指针的修改，所以后面的 g 已经发生了改变
	go modifyUser(g)
	time.Sleep(5 * time.Second)
	fmt.Println(g)
	fmt.Println(u)

	c <- g
	go printUser(c)
	fmt.Println("end")
}
