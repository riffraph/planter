package main

import "fmt"
import "io/ioutil"
import "os"
import "path/filepath"

// Args:
// inputFolder: otherwise assume current path
// outputFolder: otherwise assume current path

var fileTypes = [...]string{".pu", ".puml"} // supported file types

func main() {
	currentDir, err := os.Getwd()
	parseError(err)

	inputFolder, outputFolder := currentDir, currentDir

	args := os.Args[1:]

	if len(args) > 0 {
		inputFolder = args[0]
	}

	if len(args) > 1 {
		outputFolder = args[1]
	}

	if len(args) > 2 {
		fmt.Println("additional arguments ignored:", args[2:])
	}

	fmt.Println("Reading files from:", inputFolder)
	fmt.Println("Generating to:", outputFolder)

	// find all the compatible files in the input dir
	fileList := getFileList(inputFolder)

	for _, file := range fileList {
		fmt.Println(file)
	}
}

// getFileList
//
// get the list of compatible files in the input folder path
func getFileList(inputFolder string) []os.FileInfo {
	list := make([]os.FileInfo, 0)

	fileList, err := ioutil.ReadDir(inputFolder)
	parseError(err)

	for _, file := range fileList {
		if file.IsDir() == false {
			if filepath.Ext(file.Name()) == ".pu" {
				list = append(list, file)
			}
		}
	}

	return list
}

func parseError(err error) {
	if err != nil {
		panic(err)
	}
}
