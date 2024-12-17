package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 对象作为接收者
func (v Vertex) AbsFunc() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针作为接收者
func (v Vertex) ScaleFunc(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

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
	fmt.Printf("j=%d\n", j)

	v := Vertex{3, 4}
	fmt.Println(v.AbsFunc())
	p1 := &v
	fmt.Printf("p1.AbsFunc=%f\n", p1.AbsFunc())

	v1 := Vertex{3, 4}
	v1.ScaleFunc(10.0)
	fmt.Println(v1)
	p2 := &v1
	p2.ScaleFunc(10)
	fmt.Printf("p2=%v\n", v1)

	Scale(&v, 10)
	fmt.Println(v)
	fmt.Println(Abs(v))
	// 编译错误
	// fmt.Println(Abs(p2))

}
