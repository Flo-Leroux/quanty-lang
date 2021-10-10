parser grammar QuantyParser;

options {
    tokenVocab=QuantyLexer;
}

prog : componentDef | EOF;

componentDef : COMPONENT IDEN selectionSet;

selectionSet : LBRACE tag* RBRACE;

tag : IDEN argumentList? selectionSet?;

argumentList : LPAREN argument (COMMA argument)* RPAREN;

argument : IDEN COLON STRING;
