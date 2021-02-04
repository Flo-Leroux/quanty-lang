package main

import (
	"io/ioutil"

	"github.com/alecthomas/participle/v2"

	"github.com/alecthomas/repr"

	"quanty/ast"
	"quanty/cli"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var parser, err = participle.Build(&ast.QM{})

func main() {

	cli := cli.NewCli()
	writer := NewWriter("build")
	qm := &ast.QM{}

	for _, file := range cli.Files {
		res, err := ioutil.ReadFile(file)
		check(err)

		err = parser.ParseString("", string(res), qm)
		check(err)

		repr.Println(qm, repr.Indent("  "), repr.OmitEmpty(true))

		for _, entry := range qm.Entries {
			writer.Build(entry.Component)
		}
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
