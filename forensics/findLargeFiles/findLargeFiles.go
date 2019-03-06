package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
)

type FilesFound struct {
	FullPath string
	Info     os.FileInfo
}

func main() {

	target := "/home/chris.scogin/code/python"
	var filesList []FilesFound

	filesList = getTreeRecursive(target, filesList)

	sort.SliceStable(filesList, func(i, j int) bool {
		return filesList[i].Info.Size() < filesList[j].Info.Size()
	})

	for _, file := range filesList {
		fiSize := math.Round((float64(file.Info.Size()) / 1024))

		fmt.Printf("\nFile Path: \t%v", file.FullPath)
		fmt.Printf("\nFile Name: \t%v", file.Info.Name())
		fmt.Printf("\nFile Size(KB): \t%v\n", fiSize)
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
				FullPath: fullPath,
				Info:     dir,
			}

			list = append(list, *item)

			

		}
		// return list
	}
	return list
}
