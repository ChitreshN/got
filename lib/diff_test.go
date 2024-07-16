package lib

import (
	"io"
	"os"
	"testing"
)

func TestEditString(t *testing.T) {
	tests := []struct {
		name     string
		editList []Edit
		expected string
	}{
		{
			name: "Append test",
			editList: []Edit{
				{Append: "Hello", EditType: Append},
			},
			expected: "a5;Hello",
		},
		{
			name: "Delete test",
			editList: []Edit{
				{Delete: "World", EditType: Delete},
			},
			expected: "d5;World",
		},
		{
			name: "Identical test",
			editList: []Edit{
				{Identical: 1, EditType: Identical},
			},
			expected: "i1;",
		},
		{
			name: "Mixed edits test",
			editList: []Edit{
				{Append: "Hello", EditType: Append},
				{Identical: 2, EditType: Identical},
				{Delete: "World", EditType: Delete},
			},
			expected: "a5;Helloi2;d5;World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EditString(tt.editList)
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
        expected []Edit
    }{
        {
            "line1\nline2\nline3\n",
            "line1\nline3\nline4\n",
            []Edit{
                {Identical: 1, EditType: Identical},
                {Delete: "line2", EditType: Delete},
                {Identical: 3, EditType: Identical},
                {Append: "line4", EditType: Append},
            },
        },
        {
            "line1\nline2\n",
            "line1\nline2\nline3\n",
            []Edit{
                {Identical: 1, EditType: Identical},
                {Identical: 2, EditType: Identical},
                {Append: "line3", EditType: Append},
            },
        },
        {
            "line1\nline2\nline3\n",
            "line1\nline3\n",
            []Edit{
                {Identical: 1, EditType: Identical},
                {Delete: "line2", EditType: Delete},
                {Identical: 3, EditType: Identical},
            },
        },
    }

    for _, c := range cases {
        file1 := createTempFile(t, c.content1)
        defer os.Remove(file1.Name())
        file2 := createTempFile(t, c.content2)
        defer os.Remove(file2.Name())

        got := Diff(file1, file2)
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
