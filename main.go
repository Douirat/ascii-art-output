package main

import (
	"fmt"
	"os"
	files "ascii_art/ASCII"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("The number of arguments is not correct!!!")
		return
	} else {
		files.GenerateOutput(args[0], args[1], args[2])
	}
}
