package main

import "fmt"

func main() {
	// 两种初始化方式：
	map1 := map[string]int{"one": 1, "two": 2}
	map2 := make(map[string]string)
	map2["hello"] = "world"
	fmt.Println(map1)
	fmt.Println(map2)
	// map的key可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float32(64)
	map3 := make(map[string][]string)
	map3["test"] = []string{"map", "value"}
	map3["syke"] = []string{"lei", "ying"}
	fmt.Println(map3)
	// 判断一个key是否存在
	key := "one"
	val1, isPresent := map1[key]
	if isPresent {
		fmt.Printf("val1=%d\n", val1)
	} else {
		fmt.Printf("key %v is not exist", key)
	}
	// 遍历map
	for key, value := range map3 {
		fmt.Printf("key=%v, value=%v\n", key, value)
	}
}
