package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"time"
	"os"
)

type FilesFound struct {
	FullPath  string
	Name      string
	TimeStamp time.Time
}

func main() {

	target := os.Args[1]

	var filesList []FilesFound

	filesList = getTreeRecursive(target, filesList)

	sort.SliceStable(filesList, func(i, j int) bool {
		return filesList[i].TimeStamp.After(filesList[j].TimeStamp)
	})

	for _, file := range filesList {

		fmt.Printf("\nPath: \t\t%v", file.FullPath)
		fmt.Printf("\nName: \t\t%v", file.Name)
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
				Name:      dir.Name(),
				TimeStamp: dir.ModTime(),
			}

			list = append(list, *item)

		}
		// return list
	}
	return list
}
