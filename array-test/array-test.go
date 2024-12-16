package main

import "fmt"

func main() {
	arr := [3]string{"12312", "yingying", "syske"}
	fmt.Println(arr)
	fmt.Println(len(arr), cap(arr))
	arr1 := [3]string{}
	fmt.Println(arr1)
	fmt.Println(len(arr1), cap(arr1))

	// 切片
	slip1 := arr[1:]
	fmt.Println(slip1)
	fmt.Println(len(slip1))
	fmt.Println(cap(slip1))
	fmt.Println(cap(arr))
	slip1 = slip1[0:1]
	fmt.Println(slip1)

	fmt.Println(len(slip1))
	// 如下写法会报错：invalid argument: index -1 (constant of type int) must not be negative
	// fmt.Println(arr[:-1])
	fmt.Println(arr[:len(arr)-1])

}
