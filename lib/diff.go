package lib

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type EditType string

type Pair struct {
    first int
    second int
}

type Edit struct {
	EditType  EditType
	Append    string
	Delete    string
	Identical Pair
}

func EditString(editList []Edit) string{
    editString := ""
	for _, val := range editList{
		switch val.EditType {
		case Append:
			editString += "a"+fmt.Sprint(len(val.Append))+";"+val.Append
		case Delete:
			editString += "d"+fmt.Sprint(len(val.Delete))+";"+val.Delete
		case Identical:
			editString += "i"+fmt.Sprint(val.Identical.first)+";"+fmt.Sprint(val.Identical.second)+";"
		}
	}
    return editString
}

const (
	Append    EditType = "Append"
	Delete    EditType = "Delete"
	Identical EditType = "Identical"
)

func Diff(file1 *os.File, file2 *os.File) (editList []Edit) {
	list1 := make([]string, 0)
	list2 := make([]string, 0)

	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		list1 = append(list1, scanner.Text())
	}
	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		list2 = append(list2, scanner.Text())
	}

	editList = diff(list1, list2)
	return
}

func diff(file1 []string, file2 []string) []Edit {
	dag, _ := lcs(file1, file2)
	i := len(file1)
	j := len(file2)
	edits := make([]Edit, 0)
	for {
		if i == 0 && j == 0 {
			break
		}
		if i != 0 && dag[i][j] == dag[i-1][j] {
			i = i - 1
			edits = append(edits, Edit{Delete: file1[i], EditType: Delete})
			continue
		}
		if j != 0 && dag[i][j] == dag[i][j-1] {
			j = j - 1
			edits = append(edits, Edit{Append: file2[j], EditType: Append})
			continue
		}
		if dag[i][j] == 1+dag[i-1][j-1] {
			i, j = i-1, j-1
			edits = append(edits, Edit{Identical: Pair{i+1,j+1}, EditType: Identical})
		}
	}
	for i2, j2 := 0, len(edits)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
		edits[i2], edits[j2] = edits[j2], edits[i2]
	}
	return edits
}

func lcs(file1 []string, file2 []string) ([][]int, int) {
	file1Length := len(file1) + 1
	file2Length := len(file2) + 1
	seqArray := make([][]int, file1Length)
	for i := 0; i < file1Length; i++ {
		seqArray[i] = make([]int, file2Length)
	}
	for i := 1; i < file1Length; i++ {
		for j := 1; j < file2Length; j++ {
			lcs_helper(file1, file2, i, j, seqArray)
		}
	}
	lcsLength := seqArray[file1Length-1][file2Length-1]
	editDistance := file1Length - lcsLength + file2Length - lcsLength - 2
	return seqArray, editDistance
}

func lcs_helper(str1, str2 []string, i, j int, seqArray [][]int) {
	if str1[i-1] == str2[j-1] {
		seqArray[i][j] = 1 + seqArray[i-1][j-1]
		return
	}
	seqArray[i][j] = int(math.Max(float64(seqArray[i-1][j]), float64(seqArray[i][j-1])))
}
