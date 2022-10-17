
/** Taken from "The Definitive ANTLR 4 Reference" by Terence Parr */

// Derived from http://json.org
parser grammar QuantyParser;

options { tokenVocab=QuantyLexer; }

schema
   : components=component+ EOF
   ;

component
   : COMPONENT name=IDENT LBRACE selection+ RBRACE
   ;

selection
    : STRING | field
    ;

field
    : name=IDENT LBRACE selection+ RBRACE
    | name=IDENT
    ;
