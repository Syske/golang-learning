package main

import "fmt"

func main() {
	a := [6]int{2, 4, 5, 7, 9, 11}
	// 第一种，for-i，和其他语言基本上没有区别
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 第二种方式，省略初始化和后续逻辑
	i := 1
	for i < len(a) {
		fmt.Printf("i=%d, for a[i]=%d\n", i, a[i])
		i++
	}

	// 第三种方式，类似其他语言的while
	work := true
	j := 0
	for work {
		if j >= len(a) {
			work = false
		} else {
			fmt.Println(a[j])
		}

		j++
	}
	// 第四种方式，死循环
	fmt.Println(i)
	for {
		if i == len(a) {
			break
		}
		fmt.Println(i)
		i++
	}
}
