package main

import (
	"fmt"
	"log"
	"os"
	"math"
)

func main() {
	getFileInfo("/home/chris.scogin/.bashrc")
}

func getFileInfo(file string) {

	// get the status of the file first
	fi, err := os.Stat(file)
	if err != nil {
		log.Printf("\nError getting status of file %v: %v", file, err)
	}

	fs := (float64(fi.Size())/1024)

	fmt.Printf("\nName: \t\t\"%v\"\n", fi.Name())
	fmt.Printf("Size(KB): \t%v\n", math.Round(fs))
	fmt.Printf("Permissions: \t%v\n", fi.Mode())
	fmt.Printf("Modified Time: \t%v\n", fi.ModTime())
	if !fi.IsDir() {
		fmt.Printf("File: \t\ttrue\n")
	}
	if fi.IsDir() {
		fmt.Printf("File: \t\tfalse\n")
	}
	fmt.Printf("Directory: \t%v\n", fi.IsDir())

}
