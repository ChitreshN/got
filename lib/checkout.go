package lib

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func Checkout(commit string) error {
	cntFile, err := os.OpenFile(".got/com/cf", os.O_RDONLY, 0666)
    defer cntFile.Close()
    cur,err := GetNthline(cntFile, 2) 
    Check(err)
    for cur != commit {
        return constPrevious()
    } 
    return Revert()
}

func constNext() error {
	cntFile, err := os.OpenFile(".got/com/cf", os.O_RDONLY, 0666)
    defer cntFile.Close()
	Check(err)
    cur, _ := GetNthline(cntFile, 2)
    tot, _ := GetNthline(cntFile, 1)
    curCommit ,err := strconv.Atoi(cur)
    if err != nil{
        err = fmt.Errorf("Not an int: %s",cur)
        return err
    }

    next := curCommit + 1
    
	fileList := GetAllFiles(path.Join(".got", "com", cur))

    for _, fileName := range fileList {
		fileName,err = filepath.Rel((".got/com/"+cur),fileName)
        if err != nil {
            fmt.Printf("stupid pathh!!!: %s\n", ".got/com/"+cur)
            return err
        }
		comfName := GetComFilePath(fileName, cur)
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
		data, err := constnextCommit(latestCommit, commitInfo)
		if err != nil {
			fmt.Printf("couldnt construct commit: %v\n", err)
			return err
		}
		err = os.WriteFile(GetObjFilePath(fileName), data, 0666)
		if err != nil {
			fmt.Printf("coulndt write to file: %v\n", fileName)
		}
    }
    cf_data := tot + "\n" + fmt.Sprint(next) + "\n"
    err = os.WriteFile(".got/com/cf",[]byte(cf_data),0666)
    Check(err)
    return nil
}

func constPrevious() error {
    // assumptions - the cur commit in obj, write to obj after construction
    // so that can be called repeatedly to contruct previous commits
	cntFile, err := os.OpenFile(".got/com/cf", os.O_RDONLY, 0666)
    defer cntFile.Close()
	Check(err)
    cur, _ := GetNthline(cntFile, 2)
    tot, _ := GetNthline(cntFile, 1)
    curCommit ,err := strconv.Atoi(cur)
    if err != nil{
        err = fmt.Errorf("Not an int: %s",cur)
        return err
    }

    prev := curCommit - 1
    
	fileList := GetAllFiles(path.Join(".got", "com", cur))

    for _, fileName := range fileList {
		fileName,err = filepath.Rel((".got/com/"+cur),fileName)
        if err != nil {
            fmt.Printf("stupid pathh!!!: %s\n", ".got/com/"+cur)
            return err
        }
		comfName := GetComFilePath(fileName, cur)
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
		err = os.WriteFile(GetObjFilePath(fileName), data, 0666)
		if err != nil {
			fmt.Printf("coulndt write to file: %v\n", fileName)
		}
    }
    cf_data := tot + "\n" + fmt.Sprint(prev) + "\n"
    err = os.WriteFile(".got/com/cf",[]byte(cf_data),0666)
    Check(err)
    return nil
}
