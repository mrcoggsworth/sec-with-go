package main

import (
	"log"
	"os"
)

func main() {
	createFile("test123.txt")
}

func createFile(s string) {
	file, err := os.Create(s)
	defer file.Close()

	if err != nil {
		log.Println("CREATE FILE ERROR...", err)
	}
}
