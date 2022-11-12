package basic

import "fmt"

func Variables() {
	var a = "variable"
	fmt.Println(a) // variable

	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	var d = true
	fmt.Println(d) // true

	var e int
	fmt.Println(e) // 0

	f := "apple"
	fmt.Println(f) // apple
}