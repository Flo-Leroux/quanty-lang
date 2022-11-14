/** Taken from "The Definitive ANTLR 4 Reference" by Terence Parr */

// Derived from http://json.org
parser grammar QuantyParser;

options {
	tokenVocab = QuantyLexer;
}

schema: components = component+ EOF;

component: COMPONENT name=componentName LBRACE selection+ RBRACE;
componentName : IDENT;

selection: STRING | field;

field: name = fieldName LBRACE selection+ RBRACE | name = fieldName;
fieldName : IDENT;
