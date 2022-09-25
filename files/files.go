package files

import (
	"bufio"
	"fmt"
	"os"
)

const f_options =  os.O_APPEND | os.O_WRONLY | os.O_CREATE
const file_mode = 0644

func Write(file string, write string) {
	f, err := os.OpenFile(file, f_options, file_mode)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(f, write)
	defer f.Close()
}

func Read(file string) {
	f, err := os.OpenFile(file, os.O_RDONLY, file_mode)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	defer f.Close()
} 