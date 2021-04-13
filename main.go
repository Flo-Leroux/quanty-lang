package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/alecthomas/participle/v2"

	"quanty/ast"
	"quanty/cli"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var parser, err = participle.Build(&ast.File{})

func main() {

	cli := cli.NewCli()
	qm := &ast.File{}

	for _, file := range cli.Files {
		res, err := ioutil.ReadFile(file)
		check(err)

		err = parser.ParseString("", string(res), qm)
		check(err)

		toPrint, err := json.MarshalIndent(qm, "", "    ")
		check(err)
		fmt.Println(string(toPrint))
		// repr.Println(qm, repr.Indent("  "), repr.OmitEmpty(true))
	}

	// 	err = parser.ParseString("", `
	// package main

	// component Main {
	// 	prop name string = "test"
	// }
	// `, qm)

	// 	if err != nil {
	// 		panic(err)
	// 	}

}
