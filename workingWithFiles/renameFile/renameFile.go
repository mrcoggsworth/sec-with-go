package main

import (
	"log"
	"os"
)

func main() {
	rename("test123.txt", "test456.txt")
}

func rename(src, dst string) {
	err := os.Rename(src, dst)
	if err != nil {
		log.Println(err)
	}
}
