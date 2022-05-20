package utils

import (
	"fmt"
	"os"

	"github.com/PrinceMerluza/devcenter-content-linter/logger"
)

var (
	FileCache map[string]*string
)

func StoreFiles(files *[]string) {
	FileCache = make(map[string]*string)

	for _, file := range *files {
		// If file is already in cache, skip it
		if FileCache[file] != nil {
			continue
		}

		logger.Tracef("Opening file %s \n", file)

		// Open and store file contents in cache
		fileData, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			logger.Warnf("Error opening file %s \n: %s \n", file, err.Error())
			FileCache[file] = nil
			continue
		}

		fileString := string(fileData[:])
		FileCache[file] = &fileString
	}
}

func GetFileContent(filepath string) *string {
	return FileCache[filepath]
}
