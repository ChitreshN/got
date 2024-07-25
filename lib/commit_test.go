package lib

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestLatestCommit(t *testing.T) {
	objDir := filepath.Join(".got", "obj")
	comDir := filepath.Join(".got", "com", "1")
	err := os.MkdirAll(objDir, 0755)
	Check(err)
	err = os.MkdirAll(comDir, 0755)
	Check(err)
	err = os.WriteFile(".got/com/cf", []byte("1\n1\n"), 0666)
	Check(err)
	defer os.RemoveAll(".got")

	testCases := []struct {
		fileName       string
		fileContent    string
		commitInfo     string
		expectedResult string
	}{
		{"testfile1.txt", "line1\nline2\nline3\n", "a4;fouri2;2;i3;3;", "four\nline2\nline3\n"},
		{"testfile2.txt", "alpha\nbeta\ngamma\n", "d5;alphai2;1;a5;delta", "beta\ndelta\n"},
		{"testfile3.txt", "one\ntwo\nthree\n", "a3;ONEd3;twoi3;3;", "ONE\nthree\n"},
	}

	for _, tc := range testCases {
		objFilePath := filepath.Join(objDir, tc.fileName)
		err = os.WriteFile(objFilePath, []byte(tc.fileContent), 0666)
        if err != nil { log.Fatalln(err) }
		comFilePath := filepath.Join(comDir, tc.fileName)
		err = os.WriteFile(comFilePath, []byte(tc.commitInfo), 0666)
        if err != nil { log.Fatalln(err) }
	}

    err = LatestCommit()
	if err != nil {
		t.Fatalf("ConstLatestCommit failed: %v", err)
	}

	for _, tc := range testCases {
		reconstructedData, err := os.ReadFile(tc.fileName)
        if err != nil { log.Fatalln("errrrr",tc.fileName) }

		if string(reconstructedData) != tc.expectedResult {
			t.Errorf("For %s, expected %s, got %s", tc.fileName, tc.expectedResult, string(reconstructedData))
		}
        err = os.Remove(tc.fileName)
        if err != nil {
            t.Errorf("couldnt remove file after test: %s", tc.fileName)
        }
	}
}

