package fileTools

import (
	"fmt"
	"os"
	"path/filepath"
)

//TODO: Can this be made concurrent??
var fileList []string

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		fileList = append(fileList, path)
	}
	return nil
}

//FindFiles searches through a folder for all files matching a pattern
func FindFiles(root, pattern string) []string {
	fileList = []string{}
	fmt.Println("Walking: " + root)
	filepath.Walk(root, visit)
	var matchingFiles []string
	for _, path := range fileList {
		files, _ := filepath.Glob(path + "/" + pattern)
		fmt.Println("Files")
		fmt.Println(files)
		for _, file := range files {
			matchingFiles = append(matchingFiles, file)
			fmt.Println("Matching")
			fmt.Println(matchingFiles)
		}
	}
	return matchingFiles
}
