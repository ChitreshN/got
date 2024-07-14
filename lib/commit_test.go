package lib

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConstPrevCommit(t *testing.T) {
	// Setup test environment
	baseDir := ".got"
	err := os.MkdirAll(baseDir, 0755)
	Check(err)
	defer os.RemoveAll(baseDir) // Clean up after test

	// Create obj and com directories
	objDir := filepath.Join(baseDir, "obj")
	comDir := filepath.Join(baseDir, "com")
	err = os.MkdirAll(objDir, 0755)
	Check(err)
	err = os.MkdirAll(comDir, 0755)
	Check(err)

	// Create dummy latest commit file
	fileName := "testfile.txt"
	objFilePath := filepath.Join(objDir, fileName)
	err = os.WriteFile(objFilePath, []byte("line1\nline2\nline3\n"), 0666)
	Check(err)

	// Create dummy commit info file
	comFilePath := filepath.Join(comDir, fileName)
	err = os.WriteFile(comFilePath, []byte("a4;fouri2;i3;"), 0666)
	Check(err)

	// Call ConstPrevCommit
	err = ConstPrevCommit(fileName)
	if err != nil {
		t.Fatalf("ConstPrevCommit failed: %v", err)
	}

	// Verify the contents of the reconstructed file
	reconstructedData, err := os.ReadFile(fileName)
	Check(err)

	expectedData := "four\nline2\nline3\n"
	if string(reconstructedData) != expectedData {
		t.Errorf("Expected %s, got %s", expectedData, string(reconstructedData))
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
