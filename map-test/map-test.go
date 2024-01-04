package main

import "fmt"

func main() {
	// 两种初始化方式：
	map1 := map[string]int{"one": 1, "two": 2}
	map2 := make(map[string] string)
	map2["hello"] = "world"
	fmt.Println(map1)
	fmt.Println(map2)
}