# ALIAS
antlr4 = java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool
grun = java org.antlr.v4.gui.TestRig
antlr4-go = java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go
exportClass = export CLASSPATH=".:/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH"

# VARIABLES
mainG4 = QuantyParser
mainExprG4 = prog


all: compile

compile: clean
	${exportClass} && ${antlr4-go} ${mainG4}.g4 -o ./parser

debug: clean
	${exportClass} && ${antlr4} ${mainG4}.g4 -o ./debug
	${exportClass} && cd ./debug && javac ${mainG4}*.java
	${exportClass} && cd ./debug && ${grun} ${mainG4} ${mainExprG4} -gui

clean:
	rm -f *.java *.tokens *.interp *.class
	rm -rf ./debug
