#!/bin/sh

antlr4 -Dlanguage=Go -package parser -Xforce-atn -visitor -listener -o ./ -Xexact-output-dir ../grammar/*.g4;
# antlr4 -Dlanguage=Go -package parser -o ./grammar/atn -Xexact-output-dir -atn ./grammar/*.g4;
