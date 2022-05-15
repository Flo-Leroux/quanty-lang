package main

import (
	"os"
	"quanty/antlr/parser"

	"quanty/listener"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()
	antlr.ParseTreeWalkerDefault.Walk(listener.NewQuantyListener(), tree)

	// println("Oak Framework Initialized")

	// // Starts the Oak framework
	// oak.Start()

	// // Starts our Router
	// router.NewRouter()
	// router.RegisterRoute("", component.Home)
	// router.RegisterRoute("home", component.Home)
	// router.RegisterRoute("about", component.About)

	// // keeps our app running
	// done := make(chan struct{}, 0)
	// <-done
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"quanty/ast"
// )

// func main() {
// 	baseDoc := ast.NewDocument("Test")
// 	base, _ := json.MarshalIndent(baseDoc, "", "    ")

// 	fmt.Println(string(base))

// 	node := ast.NewComponent("Main")
// 	baseDoc.AddDefinition(node)
// 	mod, _ := json.MarshalIndent(baseDoc, "", "    ")

// 	fmt.Println(string(mod))

// 	// input, _ := antlr.NewFileStream(os.Args[1])
// 	// lexer := parser.NewQuantyLexer(input)
// 	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
// 	// p := parser.NewQuantyParser(stream)
// 	// p.BuildParseTrees = true
// 	// tree := p.Document()

// 	// var visitor = visitor.NewQuantyVisitor()
// 	// var result = visitor.Visit(tree)

// 	// res, _ := json.MarshalIndent(result, "", "   ")

// 	// fmt.Println(string(res))
// }

// package main

// //go:generate echo "Hello World!"

// import (
// 	"fmt"
// 	"quanty/listener"
// 	"quanty/parser"

// 	"github.com/antlr/antlr4/runtime/Go/antlr"
// )

// func main() {
// 	// Setup the input
// 	is, _ := antlr.NewFileStream("main.qy")

// 	// Create the Lexer
// 	lexer := parser.NewQuantyLexer(is)
// 	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

// 	p := parser.NewQuantyParser(stream)

// 	// Finally parse the expression
// 	var listener listener.QuantyListener
// 	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Document())

// 	fmt.Println(listener.Out.String())
// }
