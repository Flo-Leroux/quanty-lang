package cmd

import (
	"log"
	"os"
	"path/filepath"
)

const (
	extension = ".qy"
)

func findFiles(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if filepath.Ext(path) != extension {
			return nil
		}

		*files = append(*files, path)
		return nil
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
