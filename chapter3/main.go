package main

import "fmt"

func main() {
	var a = 8
	var b = 9
	fmt.Println("a=b结果是：", a == b)
	fmt.Println("a!=b结果是：", a != b)
	fmt.Println("a>b结果是：", a > b)
	fmt.Println("a<b结果是：", a < b)

	a1 := true
	b1 := true
	fmt.Println("a1 && b1", a1 && b1)
	fmt.Println("a1 && !b1", a1 && !b1)
	fmt.Println("a1 || b1", a1 || b1)
	fmt.Println("!a1 || b1", !a1 || b1)
	fmt.Println("!a1", !a1)
	fmt.Println("!b1", !b1)
}
