package main

import (
	"log"
	"os"
)

func main() {
	removeFile("test123.txt")
}

func removeFile(file string) {
	if err := os.Remove(file); err != nil {
		log.Println("Trouble removing the file",file,":",err)
	}

}
