package language

import (
	"path/filepath"
	"quanty/utils"
	"sync"
)

func NewTree(root string) *tree {
	tree := &tree{
		Root:  root,
		Files: make([]*File, 0),
	}
	tree.walkInRoot()

	return tree
}

type tree struct {
	Root  string
	Files []*File
}

func (t *tree) walkInRoot() {
	var wg sync.WaitGroup
	filesPath := utils.FindFiles(t.Root, ".qy")

	for _, filePath := range filesPath {
		relFilePath, err := filepath.Rel(t.Root, filePath)
		if err != nil {
			panic(err)
		}

		wg.Add(1)

		go func(filePath string) {
			defer wg.Done()
			file := NewFile(filePath)
			if file != nil {
				t.Files = append(t.Files, file)
			}
		}(relFilePath)
	}

	wg.Wait()
}
