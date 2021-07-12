// BASE ON https://github.com/vektah/gqlparser

package main

import (
	"fmt"
	"io/ioutil"
	"quanty/ast"
	. "quanty/cli"
	"quanty/parser"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	cli := NewCli()

	for _, file := range cli.Files {
		fileRead, errRead := ioutil.ReadFile(file)
		check(errRead)

		fileParsed, errParsed := parser.ParseFile(
			&ast.Source{
				Name:  file,
				Input: string(fileRead),
			},
		)

		var dumped string
		if errParsed != nil {
			dumped = ast.Dump(errParsed)
		} else {
			dumped = ast.Dump(fileParsed)
		}
		fmt.Print(dumped)
	}
}
