package main

//四种变量的声明方式
import (
	"fmt"
)

//声明全局变量，需要使用方法一，二，三的赋值方式
var gA = 100
var gB = 200

//用方法四声明全局变量 :=只能够在函数体内声明

func main() {
	//方法一：声明一个变量，默认值是0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("tpye of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcde"
	fmt.Println("bb =", bb)
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法三：在初始化时，省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Println("cc =", cc)
	fmt.Printf("type of cc = %T\n", cc)

	//方法四：(常用)省去var关键词，直接自动匹配
	d := 100
	fmt.Println("d =", d)
	fmt.Printf("type of d = %T\n", d)

	e := "hello"
	fmt.Println("e =", e)
	fmt.Printf("type of e = %T\n", e)

	f := 3.1415926
	fmt.Println("f =", f)
	fmt.Printf("type of f = %T\n", f)

	//全局变量调用
	fmt.Println("gA =", gA, "gB =", gB)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, "yy =", yy)
	var kk, ll = 100, "okok"
	fmt.Println("kk =", kk, "ll =", ll)

	//多行的声明变量
	var (
		qq = 150
		ww = true
	)
	fmt.Println("qq =", qq, "ww =", ww)
	fmt.Printf("type of ww = %T\n", ww)
}