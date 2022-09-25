package files

import "fmt"

func ReadAndWriteExample() {
	var name string
	var dirName string = "./docs/names.txt"

	fmt.Print("What is your name: ")
	fmt.Scan(&name)

	Write(dirName, name)
	Read(dirName)
}
