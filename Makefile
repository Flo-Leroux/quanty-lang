# ALIAS
antlr4 = java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool
grun = java org.antlr.v4.gui.TestRig
antlr4-go = java -Xmx500M -cp "/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go

# VARIABLES
mainG4 = Quanty
mainExprG4 = prog


all: compile

compile: clean
	cd ./parser && ${antlr4-go} ${mainG4}.g4

debug: clean
	cd ./parser && ${antlr4} ${mainG4}.g4
	cd ./parser && javac ${mainG4}*.java
	cd ./parser && ${grun} ${mainG4} ${mainExprG4} -gui

clean:
	rm -f ./parser/*.java ./parser/*.tokens ./parser/*.interp ./parser/*.go ./parser/*.class
