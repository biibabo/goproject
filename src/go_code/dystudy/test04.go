// 使用printf对齐文本，通过制定宽度
// %4v就是向左填充足够四个宽度
// 4：正数是向左填充,负数是向右填充
package main

import "fmt"

func main() {
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "junbo.zhao", 100)
}
