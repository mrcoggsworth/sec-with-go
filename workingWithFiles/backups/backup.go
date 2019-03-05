package main

import (
	"archive/tar"
	"fmt"
	"github.com/Datadog/zstd"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	bkFileName := fmt.Sprintf("practice.tar.zst")
	bkFilePath := fmt.Sprintf("/home/chris.scogin/code/%s", bkFileName)

	infiles := []string{
		"/home/chris.scogin/code",
		"/home/chris.scogin/Documents",
	}

	exfiles := []string{
		"/home/chris.scogin/code/python",
	}

	bkObj, err := os.Create(bkFilePath)
	if err != nil {
		log.Println(err)
	}

	compress(infiles, exfiles, true, bkObj)
}

func compress(include, exclude []string, verbose bool, writers ...io.Writer) error {
	mw := io.MultiWriter(writers...)

	zst := zstd.NewWriter(mw)
	defer zst.Close()

	tw := tar.NewWriter(zst)
	defer tw.Close()

	for _, src := range include {
		fileInfo, err := os.Stat(src)

		if err != nil {
			log.Println("ERROR: Unable to tar file-", src)
		}
		var baseDir string
		if fileInfo.IsDir() {
			baseDir = filepath.Base(src)
		}

		err = filepath.Walk(src, func(file string, info os.FileInfo, err error) error {
			// return on any errors
			if err != nil {
				return err
			}

			// Check to see if exclude list is empty
			if exclude != nil {
				for _, exFilePath := range exclude {
					if strings.HasPrefix(file, exFilePath) {
						return nil
					}
				}
			}

			// Verbose
			if verbose != false {
				fmt.Println(file)
			}

			newHeader, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				log.Println("HEADER ERROR:", err)
				return err
			}

			if baseDir != "" {
				newHeader.Name = filepath.Join(baseDir, strings.TrimPrefix(file, src))
			}

			// append "/" if it is a directory and not a file
			if info.IsDir() {
				newHeader.Name += "/"
			}

			if err := tw.WriteHeader(newHeader); err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			if info.IsDir() {
				return nil
			}

			f, err := os.Open(file)
			if err != nil {
				return err
			}

			if _, err := io.Copy(tw, f); err != nil {
				return err
			}

			return nil
		})

	}
	return nil
}
