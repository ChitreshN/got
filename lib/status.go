package lib

import (
	"bufio"
	"os"
	"path/filepath"
)

func getAllFiles(directory string) []string {
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
			subFiles := getAllFiles(dir)
			allFiles = append(allFiles, subFiles...)
		}
	}

	return allFiles
}

func Status() (stagedFiles, trackedFiles, untrackedFiles []string) {
	directory, err := os.Getwd()
	Check(err)
	fileList := getAllFiles(directory)

	for i := 0; i < len(fileList); i++ {
		fileList[i], err = filepath.Rel(directory, fileList[i])
		Check(err)
	}
	untrackedFiles = make([]string, 0)
	trackedFiles = make([]string, 0)

	file, err := os.OpenFile(".got/index", os.O_RDONLY, 0666)
	if err != nil {
		untrackedFiles = append(untrackedFiles, fileList...)
		file.Close()
		return
	}
	for _, fileName := range fileList {
		tracked := false
		staged := false

		idxFile, err := os.OpenFile(".got/index", os.O_RDONLY, 0666)
		Check(err)
		stgFile, err := os.OpenFile(".got/staged", os.O_RDONLY|os.O_APPEND, 0666)
		idxScanner := bufio.NewScanner(idxFile)
		stgScanner := bufio.NewScanner(stgFile)
		for stgScanner.Scan() {
			if stgScanner.Text() == fileName {
				staged = true
				stagedFiles = append(stagedFiles, fileName)
				break
			}
		}
		for idxScanner.Scan() {
			if idxScanner.Text() == fileName {
				tracked = true
				if !staged {
					trackedFiles = append(trackedFiles, fileName)
				}
				break
			}
		}
		if !tracked {
			untrackedFiles = append(untrackedFiles, fileName)
		}
		idxFile.Close()
	}
	return
}
