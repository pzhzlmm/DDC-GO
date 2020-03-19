package main

import "fmt"

func main1() {
	fmt.Println(010)  //8 八进制
	fmt.Println(0x10) //16 十六进制

	const length = 9.1
	fmt.Println(length)
	//length = length + 1 常量无法修改

	//const(){
	//	男人 = 1
	//	女人 = -1
	//	人妖 = 0
	//}
	//const 性别  = 男人
	//fmt.Println(性别)

	const mystr = "absdajho"
	fmt.Println(mystr, len(mystr)) //len求长度
}

func main() {
	//iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次
	const (
		a = iota
		b = iota
		c = iota
		d = iota
	)
	fmt.Println(a, b, c, d) //0 1 2 3
	fmt.Println(a, b, c, d) //0 1 2 3
	fmt.Println(a, b, c, d) //0 1 2 3

}
