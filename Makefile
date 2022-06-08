BINARY_NAME=quanty

# GO
GO_VERSION = 1.18.2

# ALIAS ANTLR4
antlr4 = java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool
grun = java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig
exportClass = export CLASSPATH=".:/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH"

all: clean
	compile
	link

compile:
	cd ./language/grammar && ${exportClass} && ${antlr4} ./Lexer.g4 -o ../parser -Dlanguage=Go
	cd ./language/grammar && ${exportClass} && ${antlr4} ./Parser.g4 -o ../parser -Dlanguage=Go -visitor -no-listener
	mv ./language/parser/_lexer.go ./language/parser/lexer.go
	mv ./language/parser/_parser.go ./language/parser/parser.go

link:
	go install ./cmd/quanty
	PROG=quanty source ./autocomplete/bash_autocomplete.sh

clean:
	go clean
	rm -rf ./antlr/.antlr
	rm -rf ./antlr/parser
