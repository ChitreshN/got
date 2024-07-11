package lib

import (
	"io"
	"os"
	"testing"
)

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
                {Identical: "line1", EditType: Identical},
                {Delete: "line2", EditType: Delete},
                {Identical: "line3", EditType: Identical},
                {Append: "line4", EditType: Append},
            },
        },
        {
            "line1\nline2\n",
            "line1\nline2\nline3\n",
            []Edit{
                {Identical: "line1", EditType: Identical},
                {Identical: "line2", EditType: Identical},
                {Append: "line3", EditType: Append},
            },
        },
        {
            "line1\nline2\nline3\n",
            "line1\nline3\n",
            []Edit{
                {Identical: "line1", EditType: Identical},
                {Delete: "line2", EditType: Delete},
                {Identical: "line3", EditType: Identical},
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
