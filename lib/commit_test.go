package lib

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConstPrevCommit(t *testing.T) {
	// Set up directories
	objDir := filepath.Join(".got", "obj")
	comDir := filepath.Join(".got", "com")
	err := os.MkdirAll(objDir, 0755)
	Check(err)
	err = os.MkdirAll(comDir, 0755)
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
		// Write the initial content to the obj file
		objFilePath := filepath.Join(objDir, tc.fileName)
		err = os.WriteFile(objFilePath, []byte(tc.fileContent), 0666)
		Check(err)

		// Write the commit info
		comFilePath := filepath.Join(comDir, tc.fileName)
		err = os.WriteFile(comFilePath, []byte(tc.commitInfo), 0666)
		Check(err)

		// Call ConstPrevCommit
		err = ConstPrevCommit(tc.fileName)
		if err != nil {
			t.Fatalf("ConstPrevCommit failed for %s: %v", tc.fileName, err)
		}

		// Verify the contents of the reconstructed file
		reconstructedData, err := os.ReadFile(tc.fileName)
		Check(err)

		if string(reconstructedData) != tc.expectedResult {
			t.Errorf("For %s, expected %s, got %s", tc.fileName, tc.expectedResult, string(reconstructedData))
		}

		// Clean up for the next test case
		os.Remove(tc.fileName)
	}
}
func TestGetNthline(t *testing.T) {
	// Create a temporary file

	fileName := "testfile"
	objFilePath := filepath.Join("", fileName)
    err := os.WriteFile(objFilePath, []byte("line1\nline2\nline3\n"), 0666)
	Check(err)
	// Close and reopen the file for reading
    file, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("Failed to reopen temp file: %v", err)
	}
	defer file.Close()

	// Test getting the nth line
	tests := []struct {
		lineNumber int
		expected   string
		shouldErr  bool
	}{
		{1, "line1", false},
		{2, "line2", false},
		{3, "line3", false},
		{6, "", true}, // Line 6 does not exist
	}

	for _, tt := range tests {
		// Reset file offset to the beginning for each test
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
