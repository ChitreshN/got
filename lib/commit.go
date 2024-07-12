package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Commit(message string) {
	file, err := os.OpenFile(".got/staged", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("add before commiting maybe?")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileName := scanner.Text()
        storeDiff(fileName)
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

func storeDiff(fileName string) {
    dir := filepath.Dir(fileName)
    err := os.MkdirAll(".got/com/"+dir, 0755)
    Check(err)
    objFile := ".got/obj/"+fileName
    latestCommit,err := os.OpenFile(objFile,os.O_RDONLY,0666)
    if err != nil {
        return
    }
    currentFile,err := os.OpenFile(fileName,os.O_RDONLY,0666)
    Check(err)

    commitDiff := Diff(latestCommit,currentFile)
    commitString := EditString(commitDiff)

    commitFile,err := os.OpenFile(".got/com/"+fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE,0666)
    commitFile.Write([]byte(commitString+"\n"))
    commitFile.Close()
}
