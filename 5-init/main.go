package main

//导入包若不使用，执行程序会报错
//若导入包不使用，可以在包名前加_，不会报错
//导入包名前加. 可以在func中直接调用API
import(
	"goproject/5-init/lib1"
	"goproject/5-init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}