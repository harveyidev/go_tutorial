package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

var out = 100

const (
	a = iota
	b = iota
	c = iota
	d = iota
)

//这里是程序入口
func main() {
	// 快捷键  ctrl + /

	fmt.Println(a, b, c, d)
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
	var b = 6
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

	ccc := 5
	bbb := func() {
		fmt.Println(ccc)
	}
	bbb()

	fmt.Println(out)

	test()

	// fmt.Printf("%T\n", n)
	// fmt.Printf("%v\n", n)
	// fmt.Printf("%b\n", n)
	// fmt.Printf("%d\n", n)
	// fmt.Printf("%o\n", n)
	// fmt.Printf("%x\n", n)
	fmt.Printf("%T %v", a, a)

	//8代表是2的8次方
	var a1 int8
	a1 = 127
	fmt.Println(a1)

	//8代表是2的8次方  最大值为255
	var a2 uint8 //无负号（正数）
	a2 = 255
	fmt.Println(a2)

	//数字转字符串
	// fmt.Sprint()
	// strconv.Itoa()

	var a3 uint8 = 35
	var info = "我的当前年龄"
	fmt.Println(info + fmt.Sprint(a3))

	//浮点类型转字符串
	// fmt.Sprint()
	var temp float64 //体温
	temp = 36.5
	info1 := "我当前的体温是"
	fmt.Println(info1 + fmt.Sprint(temp))

	var a4 byte = 'a'
	fmt.Println(a4)

	i1 := "abc" //i1的大小是1G

	k := i1  //把i1的值复制一份到k
	j := &i1 //j已经变成指针，指向了i1的地址，并不会把i1的值复制一份给到j
	//j 和 i1共用一个值，如果其中，j 和 i1有一个修改，则会相互影响

	fmt.Println(k, *j) //*j是把指针指向地址里的内容取出来

	var iFloat32 float32 = 3.14159265359
	var iFloat64 float64 = 3.14159265359
	fmt.Println(iFloat32, iFloat64)

	a5 := "abc我爱你"
	b5 := "abcwan"
	fmt.Println("a5的长度", len(a5))
	fmt.Println("b5的长度", len(b5))

	b6 := false
	fmt.Println("b6 类型长度：", unsafe.Sizeof(b6))

	h := "abcd\niop"
	fmt.Println(h)

	jj := `
	<html>
		<body>
		</body>
	</html>`
	fmt.Println(jj)

	var i7 int
	var f7 float64
	var t bool
	var s string
	fmt.Println(i7, f7, t, s)

	//单引号和双引号的区别
	//双引号代表一个字符串，字符串默认都有一个结束位，C语言"\0"，在文件中使用\r\n
	//单引号代表单个字符没有结束位，并且单引号里只能有一个字符

	var x = 'a'
	fmt.Println(x)

	var y = "y"
	fmt.Println(y)

	//字符串转整型
	height := "170"
	heightInt, _ := strconv.Atoi(height)
	res := heightInt - 10
	fmt.Println(res)
	//所有类型转string
	// fmt.Sprint()
	//string 转 bool
	var boolStr string = "true"
	resBool, _ := strconv.ParseBool(boolStr)
	fmt.Println(resBool)
	fmt.Printf("%T", resBool)

	//string 转 float
	var floatString = "38.234324324234"
	resFloat, _ := strconv.ParseFloat(floatString, 64)
	fmt.Println(resFloat)
	fmt.Printf("%T", resFloat)

}

func test() {
	fmt.Println(out)
	// fmt.Println(ccc)
}
