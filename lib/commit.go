package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func Commit(message string) {
	file, err := os.OpenFile(".got/staged", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("add before commiting maybe?")
	}

	scanner := bufio.NewScanner(file)

    cntFile, err := os.OpenFile(".got/com/cf", os.O_RDONLY, 0666)
    Check(err)
	line1, _ := GetNthline(cntFile, 1)
	cnt, err := strconv.Atoi(line1)
	Check(err)
	newcnt := fmt.Sprint(cnt + 1)
    cntFile.Close()

	for scanner.Scan() {
		fileName := scanner.Text()
		storeDiff(fileName,newcnt)
		data, err := os.ReadFile(fileName)
		Check(err)

		dir := filepath.Dir(fileName)
		err = os.MkdirAll(".got/obj/"+dir, 0755)
		Check(err)

		err = os.WriteFile(GetObjFilePath(fileName), data, 0666)

		Check(err)
	}

	file.Close()
	err = os.Remove(".got/staged")
}

func storeDiff(fileName string,newcnt string) {
    os.WriteFile(".got/com/cf", []byte(newcnt+"\n"+newcnt), 0666)

	dir := filepath.Dir(fileName)
    err := os.MkdirAll(".got/com/" + newcnt + "/" + dir, 0755)
	Check(err)
	latestCommit, err := os.OpenFile(GetObjFilePath(fileName), os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	currentFile, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	Check(err)

	commitDiff := Diff(currentFile, latestCommit)
	commitString := EditString(commitDiff)

	commitFile, err := os.OpenFile(GetComFilePath(fileName,newcnt), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	commitFile.Write([]byte(commitString + "\n"))
	commitFile.Close()
}

