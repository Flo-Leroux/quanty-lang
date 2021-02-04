package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"quanty/ast"

	"github.com/iancoleman/strcase"
)

// Writer is
type Writer struct {
	destFolder string
	dirPath    string
}

// NewWriter is
func NewWriter(DestFolder string) *Writer {
	w := &Writer{
		destFolder: DestFolder,
		dirPath:    fmt.Sprintf("./%s", string(DestFolder)),
	}

	w.cleanDestFolder()

	return w
}

func (w *Writer) cleanDestFolder() bool {
	os.RemoveAll(w.dirPath)
	err := os.Mkdir(w.dirPath, 0755)
	check(err)

	return true
}

// Build is
func (w *Writer) Build(c *ast.Component) bool {
	file := c.String()
	fileBytes := []byte(file)

	fileName := strcase.ToKebab(c.Name)
	filePath := fmt.Sprintf("%s/%s.js", w.dirPath, fileName)

	err = ioutil.WriteFile(filePath, fileBytes, 0644)
	check(err)

	return true
}
