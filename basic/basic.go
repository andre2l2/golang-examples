package basic

import "fmt"

func Example1() {
	var list [3]int = [3]int{1, 2, 3}
	fmt.Println(list)
}

func Example2() {
	type User struct {
		name string
		age int
	}

	user := []User{
		{name: "Anna", age: 30}, 
		{name: "Jose", age: 20},
		{name: "Igor", age: 10},
	}

	fmt.Println(user)
}

func Example3() {
	type User struct {
		name string
		age int
	}

	user := []User{
		{name: "Anna", age: 30}, 
		{name: "Jose", age: 20},
		{name: "Igor", age: 10},
	}

	for i := 0; i < len(user); i++ {
		fmt.Println(user[i].name)
	}
}

func Example4() {
	month := 3

	switch month {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
	}
}