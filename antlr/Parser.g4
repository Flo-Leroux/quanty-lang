parser grammar Parser;

options {
    tokenVocab=Lexer;
}

document : componentDef+ | EOF;

componentDef : COMPONENT name=IDEN variableDefList? selectionSet;

variableDefList : LPAREN variables+=variableDef (COMMA variables+=variableDef)* RPAREN;

variableDef : variable COLON dataType=IDEN;

variable : DOLLAR IDEN;

selectionSet : LBRACE (tagDef|STRING)+ RBRACE;

tagDef : IDEN argumentList? selectionSet?;

argumentList : LPAREN arguments+=argument (COMMA arguments+=argument)* RPAREN;

argument : IDEN COLON STRING;
