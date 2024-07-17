package lib

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

func ConstLatestCommit() error {
	cntFile, err := os.OpenFile(".got/com/cf", os.O_RDONLY, 0666)
	Check(err)

	cnt, _ := GetNthline(cntFile, 1)

	dirLength := len(".got/com/" + cnt + "/")

	fileList := GetAllFiles(path.Join(".got", "com", cnt))

	for _, fileName := range fileList {

		fileName = fileName[dirLength:]

		comfName := GetComFilePath(fileName, cnt)

		latestCommit, err := os.OpenFile(GetObjFilePath(fileName), os.O_RDONLY, 0666)
		if err != nil {
			fmt.Printf("obj file for: %s, doesnt exist\n", fileName)
			return err
		}

		comInfoFile, err := os.OpenFile(comfName, os.O_RDONLY, 0666)
		if err != nil {
			fmt.Printf("no previous commit for: %s\n", comfName)
			return err
		}

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
