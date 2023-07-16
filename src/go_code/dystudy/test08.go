// 使用if来做分支
// package main

// import "fmt"

// func main() {
// 	var command = "go east"
// 	command = "go inside"
// 	if command == "go east"{
// 		fmt.Println("You head further up the mountain.")
// 	} else if command == "go inside" {
// 		fmt.Println("You enter the cave where you live out the rest of your life.")
// 	} else {
// 		fmt.Println("Didn't quite get that.")
// 	}
// }
// -----------------------------------------------------------------------
// 逻辑运算符
// ||表示或，&&表示与 ，用于检查多个条件

// package main

// import "fmt"

// func main() {
// 	fmt.Println("The year is 2100, should you leap?")
// 	var year = 2100
// 	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
// 	if leap {
// 		fmt.Println("Look before you leap!")
// 	} else {
// 		fmt.Println("Keep your feet on the ground.")
// 	}
// }
// -----------------------------------------------------------------------
// 短路逻辑：||第一个条件是true第二个条件就不执行了
// -----------------------------------------------------------------------
// 取反逻辑运算符
// ！ 可以把true变为false，反之亦然
// package main

// import "fmt"

// func main() {
// 	var haveTorch = true
// 	var litTorch = false
// 	if !haveTorch && !litTorch {
// 		fmt.Println("Nothing to see here.")
// 	} else {
// 		fmt.Println("sweet baby")
// 	}
// }
// -----------------------------------------------------------------------
// switch
// package main

// import "fmt"

// func main() {
// 	fmt.Println("There is a cavern entrance here and a path to the east.")
// 	var command = "go inside"
// 	switch command {
// 	case "go ease":
// 		fmt.Println("You head further up the mountain.")
// 	case "enter cave", "go inside":
// 		fmt.Println("You find youeself in a dimly lit cavern")
// 	case "read sigh":
// 		fmt.Println("The sigh reads 'No Minors'.")
// 	default:
// 		fmt.Println("Didn't quite get that.")

// 	}
// }

// -----------------------------------------------------------------------

// fallthrough（坠落？贯穿？）关键字，用来执行下一个case的body部分

// package main

// import "fmt"

// func main() {
// 	var room = "lake"
// 	switch {
// 	case room == "cave":
// 		fmt.Println("You find youeself in a dimly lit cavern")
// 	case room == "lake":
// 		fmt.Println("The ice seems solid enough.")
// 		fallthrough
// 	case room == "underwater":
// 		fmt.Println("The water is freezing cold.")

// 	}
// }

// -----------------------------------------------------------------------

// 使用循环做重复
// for关键字可以让你的代码重复执行
// for后面没有跟条件就是无限循环，可以使用break跳出循环

// package main

// import (
// 	"fmt"
// 	// "time"
// )

// func main() {
// 	var count = 10
// 	// for count > 0 {
// 	for {
// 		if count < 5 {
// 			break
// 		}
// 		fmt.Println(count)
// 		// time.Sleep(time.Second)
// 		count--
// 	}
// 	fmt.Println("Littoff!")
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	const mynum = 56
// 	var num = rand.Intn(101)
// 	fmt.Println(num)
// 	for {
// 		if num == mynum {
// 			break
// 		}
// 		fmt.Println(num)
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("Please input your guess")
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input.Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		//逻辑判断
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number.please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break //猜对即退出
		}
	}
}
