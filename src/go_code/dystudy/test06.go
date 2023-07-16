// 猜数
// 使用rand包，可以生成伪随机数
// 例如，Intn可以返回一个指定范围的随机整数
// import的路径是“math/rand”
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}

// 上面的猜数好像有问题，一直是同一个数，暂时不知道为啥
