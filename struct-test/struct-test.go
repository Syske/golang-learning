package main

import "fmt"

type Struct1 struct {
	name string
	age  int
}

type Struct2 struct {
	a Struct1
}

func main() {
	// 对象初始化方式1，使用new()
	s1 := new(Struct1)
	fmt.Printf("name=%v, age=%v\n", s1.name, s1.age)

	// 初始化并设置默认值
	s2 := Struct1{"ying", 20}
	fmt.Printf("name2=%v, age2=%v\n", s2.name, s2.age)

	s3 := Struct1{age: 27, name: "ying"}
	fmt.Printf("age3=%v, name3=%v\n", s3.age, s3.name)

	s4 := Struct1{age: 30}
	fmt.Printf("age4=%v\n", s4.age)
}
