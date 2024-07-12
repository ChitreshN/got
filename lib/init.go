package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Init() {
	err := os.Mkdir(".got", 0775)
	if err != nil {
		println("Already initialized")
		os.Exit(1)
	}
	err = os.Mkdir(".got/obj", 0755)
	Check(err)
	println("Initialized empty got directory")
}

func Commit(message string) {
	file, err := os.OpenFile(".got/staged", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("add before commiting maybe?")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileName := scanner.Text()
		data, err := os.ReadFile(fileName)
		Check(err)

		dir := filepath.Dir(fileName)
		err = os.MkdirAll(".got/obj/"+dir, 0755)
		Check(err)

		objFile := ".got/obj/" + fileName
		err = os.WriteFile(objFile, data, 0666)

		Check(err)
	}

	file.Close()
	err = os.Remove(".got/staged")
}

func Add(fileName string) {
	staged, err := os.OpenFile(".got/staged", os.O_RDONLY|os.O_CREATE, 0666)

    stgScanner := bufio.NewScanner(staged)

	for stgScanner.Scan() {
		if stgScanner.Text() == fileName {
			println("File already staged")
			return
		}
	}
	staged.Close()

    index, err := os.OpenFile(".got/index", os.O_RDONLY|os.O_CREATE, 0666)
    Check(err)
    idxScanner := bufio.NewScanner(index)

	for idxScanner.Scan() {
		if idxScanner.Text() == fileName {
            staged, err = os.OpenFile(".got/staged",os.O_WRONLY|os.O_APPEND,0666)
            fileName = fileName + "\n"
            staged.Write([]byte(fileName))
            staged.Close()
            return
		}
	}
    index.Close()

	index, err = os.OpenFile(".got/index", os.O_APPEND|os.O_WRONLY, 0666)
	Check(err)
    staged, err = os.OpenFile(".got/staged",os.O_WRONLY|os.O_APPEND,0666)
	Check(err)

	fileName = fileName + "\n"
	_, err = index.Write([]byte(fileName))
    Check(err)
	_, err = staged.Write([]byte(fileName))
    Check(err)
    staged.Close()
	index.Close()
}

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
