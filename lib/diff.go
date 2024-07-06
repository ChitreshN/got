package lib

import "math"

/*
* read paper [ optimizations -> deferred to later ]
* implementation [DONE]
 */

func Diff(file1 []string, file2[]string) {
    dag,_ := Lcs(file1,file2)
    j := len(file1)
    i := len(file2)
    for {
        if i == 0 && j == 0 {
            break
        }
        if i != 0 && dag[i][j] == dag[i-1][j] {
            i = i-1
            defer println("d:",file1[i])
            continue
        }
        if j != 0 && dag[i][j] == dag[i][j-1] {
            j = j-1
            defer println("a:",file2[j])
            continue
        }
        if dag[i][j] == 1 + dag[i-1][j-1]{
            i,j = i-1,j-1
            defer println("no change:",file1[i])
        }
    }
}

func Lcs(file1 []string, file2 []string) ([][]int, int){
    file1Length := len(file1)+1
    file2Length := len(file2)+1
    seqArray := make([][]int, file1Length)
    for i :=0; i < file1Length; i++{
        seqArray[i] = make([]int, file2Length)
    }
    for i := 1; i < file1Length ; i ++ {
        for j := 1; j < file2Length; j ++ {
            lcs_helper(file1, file2, i, j, seqArray)
        }
    }
    lcsLength := seqArray[file1Length-1][file2Length-1]
    editDistance := file1Length - lcsLength + file2Length - lcsLength - 2
    return seqArray, editDistance
}

func lcs_helper(str1 , str2 []string, i, j int, seqArray [][]int) {
    if str1[i-1] == str2[j-1] {
        seqArray[i][j] = 1 + seqArray[i-1][j-1]
        return
    }
    seqArray[i][j] =  int(math.Max(float64(seqArray[i-1][j]),float64(seqArray[i][j-1])))
}
