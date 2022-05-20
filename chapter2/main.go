package main

import "fmt"

//这里是程序入口
func main() {
	// 快捷键  ctrl + /

	/*
	*	多行注释
	*	aaaa
	*	bbbb
	 */
	// a := 1
	// fmt.Println(a)
	//1、var 变量名 变量类型
	var a int
	a = 1
	fmt.Println(a)

	//2、var 变量名 = 值
	var b = 5
	fmt.Println(b)

	//3、 变量名 := 值
	c := "hello world"
	fmt.Println(c)

	//第一种：指定变量类型、声明后若不赋值，使用默认值
	//int 的默认值是0
	var i int
	fmt.Println(i)
	i = 6 //i是数字，int
	fmt.Println(i)
	ii := "i love you" //e变量就是string类型
	fmt.Println(ii)

	//同时声明多个变量，并指定变量作用域
	//第一种方式
	aa, bb, cc := 1, 2, 3
	fmt.Println(aa, bb, cc)
	//第二种方式
	var (
		d int = 4
		e int = 5
		f int = 6
	)
	fmt.Println(d, e, f)

	//作用域

	ccc := 4
	bbb := func() {
		fmt.Println(ccc)
	}
	bbb()
}
