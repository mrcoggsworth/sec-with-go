package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	getFile("http://www.devdungeon.com/archive", "contents.html")
}

func getFile(url, file string) {
	newFile, err := os.Create(file)
	defer newFile.Close()

	if err != nil {
		log.Printf("\nERROR creating %v: %v", file, err)
	}

	response, err := http.Get(url)

	if err != nil {
		log.Printf("\nERROR getting the file from the %v: %v", url, err)
	}
	defer response.Body.Close()

	if _, err := io.Copy(os.Stdout, response.Body); err != nil {
		log.Printf("\nError copying data from the %v response to the file %v: %v", url, file, err)
	}

	if _, err := io.Copy(newFile, response.Body); err != nil {
		log.Printf("\nError copying data from the %v response to the file %v: %v", url, file, err)
	}

	fmt.Printf("\nScrapped contents from %v\n", url)

}
