package main

import(
	"fmt"
)
//函数返回单个值
func fool(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c := 100
	return c
}

//函数返回多个值，但是匿名
func fool2(a string, b int) (int,int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	return 666,777
}

//函数返回多个值，且有形参名
func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("-----fool3-----")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	
	//初始化默认值为0，作用域是fool3{}
	fmt.Println("r1 =", r1)
	fmt.Println("r2 =", r2)
	
	//给有形参名的变量赋值
	r1 = 1000
	r2 = 2000
	return
}

//函数返回多个值，且形参类型一致
func fool4(a string, b int) (r1, r2 int) {
	fmt.Println("-----fool4------")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	c := fool("abc", 555)
	fmt.Println("c =", c)

	ret1, ret2 := fool2("hahaha", 999)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = fool3("xixi", 999)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)
}
