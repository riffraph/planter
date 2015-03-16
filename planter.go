package main

import "fmt"
import "io/ioutil"
import "os"
import "os/exec"
import "path/filepath"
import "bytes"

// Args:
// inputFolder: otherwise assume current path
// outputFolder: otherwise assume current path

var fileTypes = [...]string{".pu", ".puml"} // supported file types
const plantUmlPath = "./plantuml.jar"

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

	for _, puFile := range fileList {
		filename := puFile.Name()
		extension := filepath.Ext(filename)
		pngFilename := filename[0:len(filename)-len(extension)] + ".png"

		fmt.Println("pngFilename:", pngFilename)

		// check if the .png exists
		pngFile, err := os.Stat(pngFilename)
		if err == nil {
			// do not generate the image if the source file hasn't changed
			if pngFile.ModTime().After(puFile.ModTime()) {
				continue
			}
		}

		// generate the image
		generateImage(puFile)
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

// generateFiles
//
func generateImage(file os.FileInfo) {
	//cmd := exec.Command("echo", file.Name())
	cmd := exec.Command("java", "-jar", plantUmlPath, file.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	parseError(err)

	fmt.Println(out.String())
}

func parseError(err error) {
	if err != nil {
		panic(err)
	}
}
