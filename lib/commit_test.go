package lib

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestConstLatestCommit(t *testing.T) {
	// Set up directories
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
		{"testfile1.txt", "line1\nline2\nline3\n", "a4;fouri2;i3;", "four\nline2\nline3\n"},
		{"testfile2.txt", "alpha\nbeta\ngamma\n", "d5;alphai2;a5;delta", "beta\ndelta\n"},
		{"testfile3.txt", "one\ntwo\nthree\n", "a3;ONEd3;twoi3;", "ONE\nthree\n"},
	}

	for _, tc := range testCases {
		// Create object file
		objFilePath := filepath.Join(objDir, tc.fileName)
		err = os.WriteFile(objFilePath, []byte(tc.fileContent), 0666)
        if err != nil { log.Fatalln(err) }
		// Create commit info file
		comFilePath := filepath.Join(comDir, tc.fileName)
		err = os.WriteFile(comFilePath, []byte(tc.commitInfo), 0666)
        if err != nil { log.Fatalln(err) }
	}

	// Call ConstLatestCommit
	err = ConstLatestCommit()
	if err != nil {
		t.Fatalf("ConstLatestCommit failed: %v", err)
	}

	for _, tc := range testCases {
		// Verify the contents of the reconstructed file
		reconstructedData, err := os.ReadFile(tc.fileName)
        if err != nil { log.Fatalln("errrrr",tc.fileName) }

		if string(reconstructedData) != tc.expectedResult {
			t.Errorf("For %s, expected %s, got %s", tc.fileName, tc.expectedResult, string(reconstructedData))
		}

	}
}

func TestGetNthline(t *testing.T) {

	fileName := "testfile"
	objFilePath := filepath.Join("", fileName)
    err := os.WriteFile(objFilePath, []byte("line1\nline2\nline3\n"), 0666)
	Check(err)
    file, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("Failed to reopen temp file: %v", err)
	}
	defer file.Close()

	tests := []struct {
		lineNumber int
		expected   string
		shouldErr  bool
	}{
		{1, "line1", false},
		{2, "line2", false},
		{3, "line3", false},
		{6, "", true}, 
	}

	for _, tt := range tests {
		file.Seek(0, 0)

		line, err := GetNthline(file, tt.lineNumber)
		if (err != nil) != tt.shouldErr {
			t.Errorf("GetNthline(%d) error = %v, shouldErr = %v", tt.lineNumber, err, tt.shouldErr)
			continue
		}
		if line != tt.expected {
			t.Errorf("GetNthline(%d) = %v, want %v", tt.lineNumber, line, tt.expected)
		}
	}
}
