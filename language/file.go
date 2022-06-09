package language

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"quanty/language/ast"
	"quanty/language/parser"
	"quanty/utils"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewFile(path string) *File {
	file := &File{
		Path: path,
	}

	file.Document = file.toDocumentNode()

	return file
}

type File struct {
	Path     string
	Document *ast.DocumentNode
}

func (file *File) toDocumentNode() *ast.DocumentNode {

	if _, err := os.Stat(file.Path); err != nil {
		panic(err)
	}

	writeFileCache := func(path string, file *File) *ast.DocumentNode {

		doc := FileToDocumentNode(file.Path)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		}

		astContent, err := json.MarshalIndent(doc, "", "  ")
		if err != nil {
			return doc
		}

		astPath := filepath.Join(path, "qy.ast")
		astFile, err := os.Create(astPath)
		if err != nil {
			return doc
		}
		defer astFile.Close()
		astFile.Write(astContent)

		sourceContent, err := ioutil.ReadFile(file.Path)
		if err != nil {
			return doc
		}

		sourcePath := filepath.Join(path, "qy.source")
		sourceFile, err := os.Create(sourcePath)
		if err != nil {
			return doc
		}
		defer sourceFile.Close()
		sourceFile.Write(sourceContent)

		return doc
	}

	tmpRootDir := ".quanty"
	tmpDir := filepath.Join(tmpRootDir, strings.TrimSuffix(file.Path, ".qy"))
	tmpFilePath := filepath.Join(tmpDir, "qy.source")

	fileHashed := utils.HashFileContent(file.Path)
	tmpFileExist := false

	if _, err := os.Stat(tmpFilePath); err == nil {
		tmpFileExist = true
	}

	if !tmpFileExist {
		// fmt.Printf("Temp File doesn't exist for %s => Create it if possible\n", file.Path)
		return writeFileCache(tmpDir, file)
	}

	tmpFileHashed := utils.HashFileContent(tmpFilePath)

	if tmpFileHashed != fileHashed {
		// fmt.Printf("Temp File doesn't correspond for %s => Update it if possible\n", file.Path)
		return writeFileCache(tmpDir, file)
	}

	tmpAstPath := filepath.Join(tmpRootDir, tmpDir, "qy.ast")
	content, err := ioutil.ReadFile(tmpAstPath)
	if err != nil {
		// fmt.Printf("Cannot read temp File for %s => temp file is ignore\n", file.Path)
		return writeFileCache(tmpDir, file)
	}

	doc := ast.DocumentNode{}
	err = json.Unmarshal(content, &doc)
	if err != nil {
		// fmt.Printf("Cannot parse temp File for %s => temp file is ignore\n", file.Path)
		return writeFileCache(tmpDir, file)
	}

	return &doc
}

func FileToDocumentNode(filePath string) *ast.DocumentNode {
	p := newParser(filePath)
	document := p.Document().(*parser.DocumentContext)
	visitor := &visitor{}

	return visitor.VisitDocument(document)
}

func newParser(filePath string) *parser.Parser {
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	return p
}

type visitor struct {
	*parser.BaseParserVisitor
}

func (v *visitor) VisitDocument(ctx *parser.DocumentContext) *ast.DocumentNode {

	def := &ast.DocumentNode{
		Kind:        ast.DOCUMENT,
		Imports:     make([]string, 0),
		Definitions: make([]ast.DefinitionNode, 0),
	}

	def.Module = v.VisitModuleDef(ctx.ModuleDef().(*parser.ModuleDefContext))

	for _, child := range ctx.AllImportDef() {
		importDef := v.VisitImportDef(child.(*parser.ImportDefContext))
		def.Imports = append(def.Imports, importDef)
	}

	for _, child := range ctx.AllOperationDef() {
		var operationDef ast.DefinitionNode
		switch t := child.(type) {
		case *parser.ComponentDefContext:
			operationDef = v.VisitComponentDef(t)
		}

		if operationDef != nil {
			def.Definitions = append(def.Definitions, operationDef)
		}
	}

	return def
}

func (v *visitor) VisitModuleDef(ctx *parser.ModuleDefContext) string {
	return ctx.GetName().GetText()
}

func (v *visitor) VisitImportDef(ctx *parser.ImportDefContext) string {
	return ctx.GetModule().GetText()
}

func (v *visitor) VisitComponentDef(ctx *parser.ComponentDefContext) *ast.OperationDefinitionNode {

	def := &ast.OperationDefinitionNode{
		Kind:      ast.OPERATION_DEFINITION,
		Operation: ast.COMPONENT,
		Name: &ast.NameNode{
			Kind:  ast.NAME,
			Value: ctx.GetName().GetText(),
		},
		VariableDefinitions: make([]*ast.VariableDefinitionNode, 0),

		SelectionSet: v.VisitSelectionSet(ctx.SelectionSet().(*parser.SelectionSetContext)),
	}

	return def
}

// func (v *visitor) VisitVariableDefList(ctx *parser.VariableDefListContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *visitor) VisitVariableDef(ctx *parser.VariableDefContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *visitor) VisitNamedType(ctx *parser.NamedTypeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *visitor) VisitNonNullType(ctx *parser.NonNullTypeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

// func (v *visitor) VisitListType(ctx *parser.ListTypeContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func (v *visitor) VisitSelectionSet(ctx *parser.SelectionSetContext) *ast.SelectionSetNode {
	def := &ast.SelectionSetNode{
		Kind:       ast.SELECTION_SET,
		Selections: make([]ast.SelectionNode, 0),
	}

	for _, child := range ctx.AllSelection() {
		var sel ast.SelectionNode
		switch t := child.(type) {
		case *parser.SelectStringContext:
			sel = v.VisitSelectString(t)
		case *parser.SelectFieldContext:
			sel = v.VisitSelectField(t)
		}

		if sel != nil {
			def.Selections = append(def.Selections, sel)
		}
	}

	return def
}

func (v *visitor) VisitSelectString(ctx *parser.SelectStringContext) *ast.StringValueNode {

	str := ctx.GetText()

	return &ast.StringValueNode{
		Kind:  ast.STRING,
		Value: str,
	}
}

func (v *visitor) VisitSelectField(ctx *parser.SelectFieldContext) *ast.FieldNode {
	if field := ctx.Field(); field != nil {
		return v.VisitField(field.(*parser.FieldContext))
	}

	return nil
}

func (v *visitor) VisitField(ctx *parser.FieldContext) *ast.FieldNode {
	def := &ast.FieldNode{
		Kind: ast.FIELD,
		Name: &ast.NameNode{
			Kind:  ast.NAME,
			Value: ctx.GetName().GetText(),
		},
		Arguments: make([]*ast.ArgumentNode, 0),
	}

	if alias := ctx.GetAlias(); alias != nil {
		def.Alias = &ast.NameNode{
			Kind:  ast.NAME,
			Value: alias.GetText(),
		}
	}

	if ctx.ArgumentList() != nil {
		def.Arguments = v.VisitArgumentList(ctx.ArgumentList().(*parser.ArgumentListContext))
	}

	if ctx.SelectionSet() != nil {
		def.SelectionSet = v.VisitSelectionSet(ctx.SelectionSet().(*parser.SelectionSetContext))
	}

	return def
}

func (v *visitor) VisitArgumentList(ctx *parser.ArgumentListContext) []*ast.ArgumentNode {
	list := make([]*ast.ArgumentNode, 0)

	for _, child := range ctx.AllArgument() {
		var arg *ast.ArgumentNode
		switch t := child.(type) {
		case *parser.ArgumentStringContext:
			arg = v.VisitArgumentString(t)
		case *parser.ArgumentVariableContext:
			arg = v.VisitArgumentVariable(t)
		}

		if arg != nil {
			list = append(list, arg)
		}
	}

	return list
}

func (v *visitor) VisitArgumentVariable(ctx *parser.ArgumentVariableContext) *ast.ArgumentNode {
	return &ast.ArgumentNode{
		Kind: ast.ARGUMENT,
		Value: &ast.VariableNode{
			Kind: ast.VARIABLE,
			Name: &ast.NameNode{
				Kind:  ast.NAME,
				Value: ctx.GetValue().GetText(),
			},
		},
	}
}

func (v *visitor) VisitArgumentString(ctx *parser.ArgumentStringContext) *ast.ArgumentNode {
	return &ast.ArgumentNode{
		Kind: ast.ARGUMENT,
		Value: &ast.StringValueNode{
			Kind:  ast.VARIABLE,
			Value: ctx.GetValue().GetText(),
		},
	}
}
