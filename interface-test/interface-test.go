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
	fmt.Println(rect.Add())
}
