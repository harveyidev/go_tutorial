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
	fmt.Println("!!b1", !!b1)

	age := 23
	if age <= 6 {
		fmt.Println("儿童")
	}

	if age > 7 && age < 18 {
		fmt.Println("大龄儿童")
	}

	name, phone := "", ""
	age1 := 0
	fmt.Scanf("%s %s %d", &name, &phone, &age1)

}
