package main

import (
	"fmt"
	"regexp"
)

func main() {
	searchIn := "syske 1357276 456456456 8978979kkl0909090"
	pat := "\\d+"
	// 正则匹配，返回bool类型
	ok, _ := regexp.Match(pat, []byte(searchIn))
	fmt.Println(ok)

	// 还有另一种写法
	re, _ := regexp.Compile(pat)
	ok = re.Match([]byte(searchIn))
	fmt.Println(ok)

	// 提取匹配的值，默认返回第一个匹配到的值
	subtle := re.FindString(searchIn)
	fmt.Println(string(subtle))

	// 获取到匹配值的索引，后面的参数n表示匹配的数量
	findIndex := re.FindAllIndex([]byte(searchIn), 3)
	fmt.Println(findIndex)

	// 获取所有的匹配项
	findAllString := re.FindAllString(searchIn, 4)
	fmt.Println(findAllString)
}
