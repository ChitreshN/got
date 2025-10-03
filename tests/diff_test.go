package lib

import (
	"io"
	"os"
	"testing"

	"github.com/ChitreshN/got/lib"
)


type Pair struct {
    first int
    second int
}

func TestEditString(t *testing.T) {
	tests := []struct {
		name     string
		editList []lib.Edit
		expected string
	}{
		{
			name: "Append test",
			editList: []lib.Edit{
				{Append: "Hello", EditType: lib.Append},
			},
			expected: "a5;Hello",
		},
		{
			name: "Delete test",
			editList: []lib.Edit{
				{Delete: "World", EditType: lib.Delete},
			},
			expected: "d5;World",
		},
		{
			name: "Identical test",
			editList: []lib.Edit{
				{Identical: lib.Pair{First :1, Second :1}, EditType: lib.Identical},
			},
			expected: "i1;1;",
		},
		{
			name: "Mixed edits test",
			editList: []lib.Edit{
				{Append: "Hello", EditType: lib.Append},
				{Identical: lib.Pair{First: 2,Second: 2}, EditType: lib.Identical},
				{Delete: "World", EditType: lib.Delete},
			},
			expected: "a5;Helloi2;2;d5;World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lib.EditString(tt.editList)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func createTempFile(t *testing.T, content string) *os.File {
    file, err := os.CreateTemp("", "testfile")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    _, err = file.WriteString(content)
    if err != nil {
        t.Fatalf("Failed to write to temp file: %v", err)
    }
    _, err = file.Seek(0, io.SeekStart)
    if err != nil {
        t.Fatalf("Failed to seek to beginning of temp file: %v", err)
    }
    return file
}

func TestDiff(t *testing.T) {
    cases := []struct {
        content1 string
        content2 string
        expected []lib.Edit
    }{
        {
            "line1\nline2\nline3\n",
            "line1\nline3\nline4\n",
            []lib.Edit{
                {Identical: lib.Pair{First: 1,Second: 1}, EditType: lib.Identical},
                {Delete: "line2", EditType: lib.Delete},
                {Identical: lib.Pair{First: 3,Second: 2}, EditType: lib.Identical},
                {Append: "line4", EditType: lib.Append},
            },
        },
        {
            "line1\nline2\n",
            "line1\nline2\nline3\n",
            []lib.Edit{
                {Identical: lib.Pair{First: 1,Second: 1}, EditType: lib.Identical},
                {Identical: lib.Pair{First: 2,Second: 2}, EditType: lib.Identical},
                {Append: "line3", EditType: lib.Append},
            },
        },
        {
            "line1\nline2\nline3\n",
            "line1\nline3\n",
            []lib.Edit{
                {Identical: lib.Pair{First: 1,Second: 1}, EditType: lib.Identical},
                {Delete: "line2", EditType: lib.Delete},
                {Identical: lib.Pair{First: 3,Second: 2}, EditType: lib.Identical},
            },
        },
    }

    for _, c := range cases {
        file1 := createTempFile(t, c.content1)
        defer os.Remove(file1.Name())
        file2 := createTempFile(t, c.content2)
        defer os.Remove(file2.Name())

        got := lib.Diff(file1, file2)
        if len(got) != len(c.expected) {
            t.Errorf("Diff() got %v edits, expected %v edits", len(got), len(c.expected))
        }

        for i, edit := range got {
            if edit != c.expected[i] {
                t.Errorf("Diff() edit %d: got %v, expected %v", i, edit, c.expected[i])
            }
        }
    }
}
