// 常量和变量
// const：声明常量，常量的值不可以改变
// var：声明变量，想要使用变量首先需要进行声明
// package main

// import "fmt"

// func main() {
// 	const lightSpeed = 299792 //km/s
// 	var distance = 56000000   //km
// 	fmt.Println(distance/lightSpeed, "seconds")
// 	distance = 41000000
// 	fmt.Println(distance/lightSpeed, "seconds")
// }
-----------------------------------------------------------------------
// // 同时声明多个变量

// package main

// func main(){
// 	var distance = 5615454
// 	var speed = 200000   //第一种声明方式

// 	var(
// 		distance = 5615454
// 	 	speed    = 200000   //第二种声明方式

// 	)

// 	var distance, speed = 5615454, 200000  //第三种声明方式
// 	const hoursPerDay, minutesPerHour = 24, 60   //声明常量

// }
-----------------------------------------------------------------------
// 赋值运算符

// package main

// import "fmt"

// func main() {
// 	var weight = 149.0
// 	weight = weight * 0.3548
// 	weight *= 0.3548
// 	fmt.Println(weight)

// }
-----------------------------------------------------------------------
// 自增运算符
package main

import "fmt"

func main() {
	var age = 41
	age = age + 1 //第一种写法
	age += 1      //第二种写法
	age++         //第三种写法
	// ++age          //错误的写法
	fmt.Println(age)
}
