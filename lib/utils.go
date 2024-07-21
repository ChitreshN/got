package lib

import (
	"bufio"
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		fmt.Printf("error: %v", e)
		os.Exit(1)
	}
}

func RunStatus() {
	staged, tracked, untracked := Status()
	fmt.Printf("The following files are staged for commit:\n")
	for _, fileName := range staged{
		fmt.Printf("\033[32m%s\033[0m\n", fileName)
	}
	fmt.Printf("The following files are tracked but not staged for commit:\n")
	for _, fileName := range tracked {
		fmt.Printf("\033[31m%s\033[0m\n", fileName)
	}
	fmt.Printf("The following files are untracked:\n")
	for _, fileName := range untracked {
		fmt.Printf("\033[31m%s\033[0m\n", fileName)
	}
}

func RunDiff() {
	file1, err := os.OpenFile(os.Args[2], os.O_RDONLY, 0666)
    defer file1.Close()
	Check(err)
	// how handle first commit?
	objFile := ".got/obj/" + os.Args[2]
	file2, err := os.OpenFile(objFile, os.O_RDONLY, 0666)
    defer file2.Close()
	if err != nil {
		fmt.Println("commit first before diffing it")
	}
	diff := Diff(file2, file1)
	for _, val := range diff {
		switch val.EditType {
		case Append:
			fmt.Printf("\033[32m+%s\033[0m\n", val.Append)
		case Delete:
			fmt.Printf("\033[31m-%s\033[0m\n", val.Delete)
		case Identical:
			fmt.Printf("%d\n", val.Identical)
		}
	}
}

func GetNthline(file *os.File, n int) (string,error) {
    file.Seek(0,0)
    scanner := bufio.NewScanner(file)

    lineCount := 0

    for scanner.Scan() {
        lineCount++
        if lineCount == n { return scanner.Text(),nil }
    }

    if err := scanner.Err(); err != nil { return "",err }

    return "",fmt.Errorf("no line %d in file: %s",n,file.Name())
}

func GetLastline(file *os.File) (string,error) {
    file.Seek(0,0)
    scanner := bufio.NewScanner(file)

    lasLine := ""

    for scanner.Scan() {
        lasLine = scanner.Text()
    }

    if err := scanner.Err(); err != nil { return "",err }

    return lasLine, nil
}

func getLines(fileName string) ([]string,error ){
    idxFile, err := os.OpenFile(fileName,os.O_RDONLY,0666)
    defer idxFile.Close()
    if err != nil {
        err = fmt.Errorf("No file with that name in the cur dir")
        return (make([]string,0)),err
    }
    scanner := bufio.NewScanner(idxFile)
    idxFiles := make([]string,0)
    for scanner.Scan() {
        idxFiles = append(idxFiles, scanner.Text())
    }
    return idxFiles,nil
}
