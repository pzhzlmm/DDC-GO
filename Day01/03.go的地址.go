package main

import "fmt"

func main11() {
	//基本数据类型内存各自独立
	var num1 = 100
	fmt.Println(num1, &num1) //100 0xc00000a0b8
	var num2 = 100
	fmt.Println(num2, &num2) //100 0xc00000a100

	var str1 = "hello"
	var str2 = str1
	fmt.Println(str1, &str1) //hello 0xc0000501f0
	fmt.Println(str2, &str2) //hello 0xc000050200

}

func main() {
	_, x := 100, 200 //_代表占位,无法显示读取写入
	fmt.Println(x)
}
