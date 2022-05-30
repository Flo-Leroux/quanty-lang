parser grammar Parser;

options {
    tokenVocab=Lexer;
}

file : moduleDef imports=importDef* componentDef+ | EOF;

moduleDef : MODULE name=IDEN;

importDef : FROM module=IDEN IMPORT imports+=IDEN (COMMA imports+=IDEN)*;

componentDef : COMPONENT name=IDEN variableDefList? selectionSet;

variableDefList : LPAREN variables+=variableDef (COMMA variables+=variableDef)* RPAREN;

variableDef : variable COLON dataType=IDEN;

variable : DOLLAR IDEN;

selectionSet : LBRACE selection+ RBRACE;

selection : STRING # selectString
    | tagDef # selectTag;

tagDef : IDEN argumentList? selectionSet?;

argumentList : LPAREN arguments+=argument (COMMA arguments+=argument)* RPAREN;

argument : key=IDEN COLON value=STRING;
