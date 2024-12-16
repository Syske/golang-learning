package main

import "fmt"

type Interface1 interface {
	Add() int
}

type Rectangle struct {
	width  int
	height int
}

func (re *Rectangle) Add() int {
	return (re.height + re.width)
}

func main() {
	rect := Rectangle{2, 3}
	fmt.Printf("rect.height:%d\n", rect.height)
	fmt.Printf("rect add result:%d\n", rect.Add())

	rect1 := Rectangle{width: 2}
	fmt.Printf("rect1:%s", rect1)
}
