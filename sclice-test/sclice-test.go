package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Printf("primes.len:%d, primes.cap:%d", len(primes), cap(primes))

	a := primes[1:4]
	fmt.Printf("a.value: %v, a.type:%T\n", a, a)
	fmt.Printf("a.len : %d, a.cap: %d\n", len(a), cap(a))

	a = primes[1:]
	fmt.Printf("a :%v, a.len: %d, a.cap:%d\n", a, len(a), cap((a)))

	// build new sclice
	b := make([]int, 5)
	fmt.Printf("Type: %T\n", b)
	fmt.Printf("b.vale: %v, a.len: %d, a.cap: %d\n", b, len(b), cap(b))

	c := make([]int, 0, 6)
	fmt.Printf("c.vale: %v, c.len: %d, c.cap: %d\n", c, len(c), cap(c))

}
