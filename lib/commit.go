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

	for scanner.Scan() {
		fileName := scanner.Text()
		storeDiff(fileName)
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

func storeDiff(fileName string) {
	dir := filepath.Dir(fileName)
	err := os.MkdirAll(".got/com/"+dir, 0755)
	Check(err)
	latestCommit, err := os.OpenFile(GetObjFilePath(fileName), os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	currentFile, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	Check(err)

	commitDiff := Diff(currentFile, latestCommit)
	commitString := EditString(commitDiff)

	commitFile, err := os.OpenFile(GetComFilePath(fileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	commitFile.Write([]byte(commitString + "\n"))
	commitFile.Close()
}

func ConstLatestCommit(fileName string) error {
	latestCommit, err := os.OpenFile(GetObjFilePath(fileName), os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("obj file for: %s, doesnt exist\n", fileName)
		return err
	}
	comfName := GetComFilePath(fileName)

	comInfoFile, err := os.OpenFile(comfName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("no previous commit for: %s\n", comfName)
		return err
	}

	// what do here
	commitInfo, err := GetLastline(comInfoFile)

	data, err := constCommit(latestCommit, commitInfo)
	if err != nil {
		fmt.Printf("couldnt construct commit: %v\n", err)
		return err
	}

	err = os.WriteFile(fileName, data, 0666)
	if err != nil {
		fmt.Printf("coulndt write to file: %v\n", fileName)
	}
	return nil
}

func constCommit(prevCommit *os.File, commitString string) ([]byte, error) {
	data := ""
	for i := 0; i < len(commitString); {
		switch commitString[i] {
		case 'i':
			s := i
			for commitString[i] != ';' {
				i++
			}
			lineNo, err := strconv.Atoi(commitString[s+1 : i])
			if err != nil {
				fmt.Printf("skill issues: %v", err)
				return []byte(""), err
			}
			line, err := GetNthline(prevCommit, lineNo)
			if err != nil {
				return []byte(""), err
			}
			data += line + "\n"
			i += 1

		case 'a':
			s := i
			for commitString[i] != ';' {
				i++
			}
			dataLength, err := strconv.Atoi(commitString[s+1 : i])
			if err != nil {
				fmt.Printf("skill issues: %v", err)
				return []byte(""), err
			}
			data += commitString[i+1:i+1+dataLength] + "\n"
			i += dataLength + 1

		case 'd':
			s := i
			for commitString[i] != ';' {
				i++
			}
			dataLength, err := strconv.Atoi(commitString[s+1 : i])
			if err != nil {
				fmt.Printf("skill issues: %v", err)
				return []byte(""), err
			}
			i += dataLength + 1
		default:
			err := fmt.Errorf("unknown command:%s ", string(commitString[i]))
			return []byte(""), err
		}
	}
	return []byte(data), nil
}

func constnextCommit(currentCommit *os.File, commitString string) ([]byte, error) {
	data := ""
	for i := 0; i < len(commitString); {
		switch commitString[i] {
		case 'i':
			s := i
			for commitString[i] != ';' {
				i++
			}
			lineNo, err := strconv.Atoi(commitString[s+1 : i])
			if err != nil {
				fmt.Printf("skill issues: %v", err)
				return []byte(""), err
			}
			line, err := GetNthline(currentCommit, lineNo)
			if err != nil {
				return []byte(""), err
			}
			data += line + "\n"
			i += 1

		case 'd':
			s := i
			for commitString[i] != ';' {
				i++
			}
			dataLength, err := strconv.Atoi(commitString[s+1 : i])
			if err != nil {
				fmt.Printf("skill issues: %v", err)
				return []byte(""), err
			}
			data += commitString[i+1:i+1+dataLength] + "\n"
			i += dataLength + 1

		case 'a':
			s := i
			for commitString[i] != ';' {
				i++
			}
			dataLength, err := strconv.Atoi(commitString[s+1 : i])
			if err != nil {
				fmt.Printf("skill issues: %v", err)
				return []byte(""), err
			}
			i += dataLength + 1
		default:
			err := fmt.Errorf("unknown command:%s ", string(commitString[i]))
			return []byte(""), err
		}
	}
	return []byte(data), nil
}
