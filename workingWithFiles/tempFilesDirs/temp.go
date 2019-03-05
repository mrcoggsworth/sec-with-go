package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	stuff := []byte("This is a test to write to a temp file")

	tempDirName := tmpDir()
	defer os.Remove(tempDirName)

	tempFileName := tmpFile(tempDirName, "example.txt", stuff)
	defer os.Remove(tempFileName)

	fmt.Println("done...")

}

func tmpDir() string {

	td, err := ioutil.TempDir("", "/tmp")

	if err != nil {
		log.Println("Issue creating a temp directory,", err)
	}

	fmt.Println("Temporary directory", td, "has been created...")

	return td
}

func tmpFile(path, file string, contents []byte) string {
	tf, err := ioutil.TempFile(path, file)
	defer tf.Close()

	if err != nil {
		log.Println("Issue creating a temporary file", err)
	}

	if _, err = tf.Write(contents); err != nil {
		log.Println("Could not write contents to temporary file.")
	}
	return tf.Name()
}
