package basic

import (
	"fmt"
	"math"
)

const s string = "this is a constant"

func Constants() {
	fmt.Println(s) // this is a constant

	const n = 5000

	const d = 3e20 / n

	fmt.Println(int64(d)) // 60000000000000000

	fmt.Println(math.Sin(n)) // -0.9879664387667769
}