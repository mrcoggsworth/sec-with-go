package main

import (
	"log"
	"os"
)

func main() {
	truncateFile("test123.txt")
}

func truncateFile(s string) {
	err := os.Truncate(s, 10) //the second param size is in bytes

	if err != nil {
		log.Println("TRUNCATING ERROR...", err)
	}

}
