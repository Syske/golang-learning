package main

import "fmt"

func main() {
	i, j := 42, 520

	p := &i

	fmt.Println(*p)
	fmt.Println(p)
	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Printf("Type:%T, Value:%v, target:%d\n", p, p, *p)
	*p = *p / 10
	// print 21
	// fmt.Println(p)
	// print 52
	fmt.Println(j)

}
