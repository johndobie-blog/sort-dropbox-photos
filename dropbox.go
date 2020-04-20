package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var uploadDir, dateFormat string

/**
	This will process all of the files in the passed in camera uploads folder.
	If a file has a filename in the correct format ,  it will be moved into a new sub folder.

	For example, 2019-11-17 14.30.17.jpg will be moved into /2019/11
 */
func main() {

	uploadDir = *flag.String("uploadDir", "D:/Dropbox/Camera Uploads", "This is the directory that contains your DropBox Camera uploads.")
	dateFormat = *flag.String("dateFormat", "2006-01-02", "This is date format used in the UK, for example YYYY-DD-MM")
	flag.Parse()

	HandleUploads(uploadDir)
}

func HandleUploads(path string) {

	fmt.Printf("Looking for files in [%s]\n", uploadDir)
	matches, _ := filepath.Glob(uploadDir + "/*.*")
	for _, path := range matches {
		HandleSingleFile(path)
	}
}

func HandleSingleFile(path string) {

	filename := filepath.Base(path)
	safeSubstring := filename[0:10]
	t, err := time.Parse(dateFormat, safeSubstring)

	// This is a file with a filename in the correct format for moving
	if err == nil {
		MoveSingleFile(path, filename, strconv.Itoa(t.Year()), int(t.Month()))
	}
}

func MoveSingleFile(path string, filename string, year string, month int) {

	yearDirectory := uploadDir + "/" + year

	monthValue := fmt.Sprintf("%02d", month)
	monthDirectory := yearDirectory + "/" + monthValue

	targetPath := monthDirectory + "/" + filename

	if DirectoryDoesNotExist(yearDirectory) {
		fmt.Printf("The Target Directory Does Not Exist. Creating The Directory[%s] [%s]\n", yearDirectory, filename)
		CreateNewDirectory(yearDirectory)
	}

	if DirectoryDoesNotExist(monthDirectory) {
		fmt.Printf("The Target Directory Does Not Exist. Creating The Directory [%s]\n", monthDirectory)

		CreateNewDirectory(monthDirectory)
	}

	MoveFile(path, targetPath)
}

func CreateNewDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DirectoryDoesNotExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// exist
		return false
	}
	return true
}

func MoveFile(path string, targetPath string) {
	fmt.Printf("Moving File From [%s] to [%s]\n", path, targetPath)
	err := os.Rename(path, targetPath)
	if err != nil {
		log.Fatal(err)
	}
}