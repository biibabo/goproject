package main

import "fmt"

func main() {

	// 转义字符的使用
	fmt.Println("tom\tjack") //输出：tom     jack
	fmt.Println("hello\nworld")
	fmt.Println("E:\\Go_WorkSpace\\myproject\\src")
	fmt.Println("tom说:\"i love you\"")
	//r=回车，从当前行的最前面开始输出，覆盖以前的输出
	fmt.Println("太牛龙八部\r张飞")
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")

}
