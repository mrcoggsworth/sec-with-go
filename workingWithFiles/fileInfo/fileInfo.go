package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	getFileInfo("test123.txt")
}

func getFileInfo(s string) {
	file, err := os.Stat(s)

	if err != nil {
		log.Println("GET FILE INFO ERROR...", err)
	}
	fmt.Printf("\nName:\t\t%s\n", file.Name())
	fmt.Printf("Size:\t\t%v(KB)\n", file.Size())
	fmt.Printf("Permissions:\t%v\n", file.Mode())
	if file.IsDir() == false {
		fmt.Printf("File:\t\t%v\n", true)
	}
	fmt.Printf("Directory:\t%v\n", file.IsDir())
	fmt.Printf("Modified Time:\t%v\n", file.ModTime())
}
