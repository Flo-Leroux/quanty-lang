package main

import (
	"io/ioutil"

	"strings"

	"github.com/alecthomas/repr"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	cli := NewCli()

	for _, file := range cli.Files {
		res, err := ioutil.ReadFile(file)
		check(err)
		repr.Println(string(res))
		dataReaded := strings.NewReader(string(res))
		parser := NewParser(dataReaded)
		parsed, err := parser.Parse()

		if err != nil {
			repr.Println(err)
		} else {
			// transformer := NewTransformer(parsed)

			// file := ""

			// for _, compAst := range transformer.ast.Components {
			// 	file += transformer.transformComponent(compAst)
			// }

			// if err != nil {
			// 	repr.Println(err)
			// } else {
			// 	repr.Println(transformed)

			// dataToFile := []byte(file)
			// ioutil.WriteFile("./test/main.jsx", dataToFile, 777)
			// }
			repr.Println(parsed.Components)
		}
	}
}
