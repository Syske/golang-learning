package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("main")
	str := "hello"
	var buf bytes.Buffer
	buf.WriteString(str)
	buf.WriteString(" ")
	buf.WriteString("world")
	buf.WriteString("!")
	fmt.Println(buf.String())
}
