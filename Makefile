# ALIAS
antlr4 = java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool
grun = java -Xmx500M -cp "/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig
exportClass = export CLASSPATH=".:/usr/local/lib/antlr-4.10.1-complete.jar:$CLASSPATH"

all: compile

compile: clean
	cd ./language/grammar && ${exportClass} && ${antlr4} ./Lexer.g4 -o ../parser -Dlanguage=Go
# cd ./antlr && ${exportClass} && ${antlr4} ./Parser.g4 -o ./parser -Dlanguage=Go
	cd ./language/grammar && ${exportClass} && ${antlr4} ./Parser.g4 -o ../parser -Dlanguage=Go -visitor -no-listener
	mv ./language/parser/_lexer.go ./language/parser/lexer.go
	mv ./language/parser/_parser.go ./language/parser/parser.go


json:


# debug: clean
# 	${exportClass} && ${antlr4} ${name}Lexer.g4 -o ./debug
# 	${exportClass} && ${antlr4} ${name}Parser.g4 -o ./debug -lib ./debug
# 	${exportClass} && cd ./debug && javac ${name}*.java
# 	${exportClass} && cd ./debug && ${grun} ${name} ${expression} -gui

clean:
	rm -rf ./antlr/.antlr
	rm -rf ./antlr/parser
