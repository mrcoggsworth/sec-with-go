package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type FilesFound struct {
	FullPath  string
	Info      os.FileInfo
	TimeStamp time.Time
}

func main() {

	target := os.Args[1]
	var filesList []FilesFound

	filesList = getTreeRecursive(target, filesList)

	sort.SliceStable(filesList, func(i, j int) bool {
		return filesList[i].Info.Size() > filesList[j].Info.Size()
	})

	for _, file := range filesList {
		fiSizeKB := math.Round((float64(file.Info.Size()) / 1024))
		fiSizeMB := (float64(file.Info.Size()) / 1024)/1024
		fiSizeGB := (float64((file.Info.Size()) / 1024)/1024)/1024

		fmt.Printf("\nPath: \t\t%v", file.FullPath)
		fmt.Printf("\nName: \t\t%v", file.Info.Name())
		fmt.Printf("\nSize(KB): \t%v", fiSizeKB)
		fmt.Printf("\nSize(MB): \t%.2f", fiSizeMB)
		fmt.Printf("\nSize(GB): \t%.2f", fiSizeGB)
		fmt.Printf("\nTimestamp: \t%v\n", file.TimeStamp)
	}

}

func getTreeRecursive(path string, list []FilesFound) []FilesFound {

	dirTree, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, dir := range dirTree {

		fullPath := filepath.Join(path, dir.Name())

		if dir.IsDir() {
			list = getTreeRecursive(filepath.Join(path, dir.Name()), list)

		} else if dir.Mode().IsRegular() {

			item := &FilesFound{
				FullPath:  fullPath,
				Info:      dir,
				TimeStamp: dir.ModTime(),
			}

			list = append(list, *item)

		}
		// return list
	}
	return list
}
