# ALIAS
antlr4 = java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool
grun = java org.antlr.v4.gui.TestRig
exportClass = export CLASSPATH=".:/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH"

# VARIABLES
name = Quanty
expression = prog

all: compile

compile: clean
	${exportClass} && ${antlr4} ${name}Lexer.g4 -o ./parser -Dlanguage=Go
	${exportClass} && ${antlr4} ${name}Parser.g4 -o ./parser -Dlanguage=Go

debug: clean
	${exportClass} && ${antlr4} ${name}Lexer.g4 -o ./debug
	${exportClass} && ${antlr4} ${name}Parser.g4 -o ./debug -lib ./debug
	${exportClass} && cd ./debug && javac ${name}*.java
	${exportClass} && cd ./debug && ${grun} ${name} ${expression} -gui

clean:
	rm -f *.java *.tokens *.interp *.class
	rm -rf ./debug
