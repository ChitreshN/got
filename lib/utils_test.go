package lib

import (
	"os"
	"path/filepath"
	"testing"
)


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
    os.Remove(fileName)
}
