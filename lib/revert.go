package lib

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func Revert() error {
    idxFiles,err := GetIndexFiles()
    Check(err)
    
    for _,file := range idxFiles {
        data,err := os.ReadFile(GetObjFilePath(file))
        if err != nil {
            fmt.Printf("done effed up: %v",err)
            return err
        }
        err = os.WriteFile(file,data,0666)
        if err != nil {
            fmt.Printf("cant revert file: %s,err: %v",file,err) 
            return err
        }
    }
    return nil
}


func LatestCommit() error {
	cntFile, err := os.OpenFile(".got/com/cf", os.O_RDONLY, 0666)
    defer cntFile.Close()
	Check(err)
	cnt, _ := GetNthline(cntFile, 1)
	fileList := GetAllFiles(path.Join(".got", "com", cnt))

	for _, fileName := range fileList {
		fileName,err = filepath.Rel((".got/com/"+cnt),fileName)
        if err != nil {
            fmt.Printf("stupid pathh!!!: %s\n", ".got/com/"+cnt)
            return err
        }
		comfName := GetComFilePath(fileName, cnt)
        latestCommit, err := os.OpenFile(GetObjFilePath(fileName), os.O_RDONLY, 0666)
        defer latestCommit.Close()
		if err != nil {
			fmt.Printf("obj file for: %s, doesnt exist\n", fileName)
			return err
		}
		comInfoFile, err := os.OpenFile(comfName, os.O_RDONLY, 0666)
        defer comInfoFile.Close()
		if err != nil {
			fmt.Printf("no previous commit for: %s\n", comfName)
			return err
		}
		commitInfo, err := GetLastline(comInfoFile)
        if err != nil {
            fmt.Printf("last line err: %v\n",err)
            return err
        }
		data, err := constCommit(latestCommit, commitInfo)
		if err != nil {
			fmt.Printf("couldnt construct commit: %v\n", err)
			return err
		}
		err = os.WriteFile(fileName, data, 0666)
		if err != nil {
			fmt.Printf("coulndt write to file: %v\n", fileName)
		}
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
			for commitString[i] != ';' {
				i++
			}
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
            i += 1
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
