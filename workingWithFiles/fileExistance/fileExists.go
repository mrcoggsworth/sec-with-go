package main

import (
	"fmt"
	"os"
)

func main() {
	fileExists(os.Args[1])
}

func fileExists(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("No such file or directory")
	} else {
		fmt.Println("File exists...")
	}
}
