package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// JSON字符串
	jsonStr := `{ "name": "Alice", "age": 25 }`

	var person Person // 定义Person结构体变量

	err := json.Unmarshal([]byte(jsonStr), &person) // 将JSON字符串转换为struct对象
	if err != nil {
		panic(err)
	}

	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)

	personStr, _ := json.Marshal(person)

	fmt.Println(string(personStr))
}
