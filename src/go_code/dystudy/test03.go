// 格式化打印
// 可以使用 Printf 来控制打印的输出结果。
// 与 Print 和 Println 不同，Printf的第一个参数必须是字符串。
// 这个字符串里包含了像 %v 这样的格式化动词，它的值由第二个参数的值所代替。
// 如果指定了多个格式化动词，那么它们的值由后边的参数值按其顺序进行替换。
package main

import "fmt"

func main() {
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v  years old.\n", 41*365/687)

	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)
}
