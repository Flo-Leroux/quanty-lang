package language

import (
	"log"
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
		wg.Add(1)

		go func(filePath string) {
			defer wg.Done()
			file := NewFile(filePath)
			log.Println(filePath, file != nil)
			if file != nil {
				t.Files = append(t.Files, file)
			}
		}(filePath)
	}

	wg.Wait()
}
