package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	copy(os.Args[1], os.Args[2])
}

func copy(src, dst string) {
	ogFile, err := os.Open(src)
	if err != nil {
		log.Println(err)
	}

	defer ogFile.Close()

	newFile, err := os.Create(dst)
	if err != nil {
		log.Println(err)
	}

	defer newFile.Close()

	if _, err := io.Copy(newFile, ogFile); err != nil {
		log.Println(err)
	}

	fileInfo, err := os.Stat(dst)
	if err != nil {
		log.Println(err)
	}

	if !os.IsNotExist(err) {
		fmt.Printf("\nFile %v has been copied to %v\n", src, fileInfo.Name())
	}
}
