package main

import (
	"fmt"
)

type AddInterface interface {
	Add() int
	Scale(int) int
}

type Rectangle struct {
	width  int
	height int
}

func (re *Rectangle) Add() int {
	return (re.height + re.width)
}

func (r *Rectangle) Scale(size int) int {
	return (r.height + r.width) * size
}

func main() {

	var rect AddInterface = &Rectangle{width: 2, height: 3}
	// var rect1 AddInterface = &Rectangle{2, 3}
	// fmt.Printf("rect.height:%d\n", rect.Height)
	fmt.Printf("rect add result:%d\n", rect.Add())

	rect1 := Rectangle{width: 2}
	fmt.Printf("rect1:%v", rect1)
}
