// Boolean类型
// 如果xxx，那么xxx；如果xxx是真，那么请执行xxx（true和false）
// strings.Contains:来自strings包的Contains函数可以判断某个
// 字符串是否包含另外要给的字符串
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	fmt.Println("You find yourself in a dimly lit cavrn.")
// 	var command = "walk outside"
// 	var exit = strings.Contains(command, "outside") // 判断command是否包含outside
// 	fmt.Println("You leave the cave:", exit)

// }
-----------------------------------------------------------------------
// 比较
// 如果我们比较两个值，得到的结果是true和false
// 比较运算符：==, <=, <, !=, >=, >

package main

import (
	"fmt"
)

func main() {
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
}
