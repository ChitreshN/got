package lib

import (
	"os"
	"path"
	"path/filepath"
)

func GetObjFilePath(file string) string {
	return path.Join(".got", "obj", file)
}

func GetComFilePath(file string, dir string) string {
	return path.Join(".got", "com", dir, file)
}

func GetAllFiles(directory string) []string {
	var allFiles []string
	fileList, err := os.ReadDir(directory)
	Check(err)

	for _, fileInfo := range fileList {
		if fileInfo.Name() == ".git" || fileInfo.Name() == ".got" {
			continue
		}
		if !fileInfo.IsDir() {
			allFiles = append(allFiles, filepath.Join(directory, fileInfo.Name()))
		} else {
			dir := filepath.Join(directory, fileInfo.Name())
			subFiles := GetAllFiles(dir)
			allFiles = append(allFiles, subFiles...)
		}
	}

	return allFiles
}
