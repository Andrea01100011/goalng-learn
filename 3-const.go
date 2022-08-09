package main

import "fmt"

//const定义枚举类型
const (
	//可以在const()中添加关键字iota，每行的iota累加1，第一行的iota默认为0
	BEIJING   = iota //iota = 0
	SHANGHAI         //iota = 1
	GUANGZHOU        //iota = 2
	SHENZHEN         //iota = 3
)
const (
	a, b = iota + 1, iota + 2 //iota=0, a=iota+1=0+1=1
	c, d                      //iota=1, c=iota+1=2
	e, f                      //iota=2, e=iota+1=3
	g, h = iota * 2, iota * 3 //iota=3, g=iota*2=6
	i, j                      //iota=4, i=iota*2=8
)

func main() {
	//常量（只读属性）不允许赋值
	const length int = 10
	fmt.Println("length =", length)

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)

	fmt.Println("a =", a, "b =", b, "c =", c, "d =", d)
	fmt.Println("g =", g, "h =", h, "i =", i, "j =", j)
}
