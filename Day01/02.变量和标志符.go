package main

import "fmt"

func main9() {
	var 张宏杰 = "逗比"
	fmt.Println(张宏杰)
}

func main2() {
	var num int
	num = 10
	fmt.Print("女神有", num, "备胎")
	var length = 8.1
	fmt.Print("林xxyo", length)
}

func main3() {
	var 我的名字 string = "momo"
	var 他的年龄 int = 15 //是否指定类型其实都可以
	var 长度 = 13.5
	fmt.Print(我的名字, 他的年龄, 长度)
}
func main4() { //另外一种赋值方法
	c := 10
	fmt.Print(c)
}
func main() {
	//多个变量一起赋值
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
	var x, y, z = 4, 6, 7
	fmt.Println(x, y, z)
	//简单方法
	t1, t2, t3 := 10, 11, 12
	fmt.Println(t1, t2, t3)
}
