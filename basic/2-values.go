package basic

import (
	"fmt"
)

func Values() {
	fmt.Println("go" + "lang") // go lang

	fmt.Println("1 + 1 =", 1+1) // 2
	fmt.Println("7.0 / 3.0", 7.0 / 3.0) // 2.33333333333333352.3333333333333335

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true) // false
}