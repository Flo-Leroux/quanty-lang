parser grammar Parser;

options {
    tokenVocab=Lexer;
}

document : moduleDef imports+=importDef* operations+=operationDef*
    | EOF;

moduleDef : MODULE name=IDEN;

importDef : FROM module=IDEN IMPORT imports+=IDEN (COMMA imports+=IDEN)*;

operationDef :
    COMPONENT name=IDEN (variableDefList)? selectionSet #componentDef;

variableDefList : PAREN_L variables+=variableDef (COMMA variables+=variableDef)* PAREN_R;

variableDef : DOLLAR variable=IDEN COLON type (EQUALS defaultValue=STRING)?;

type : BRACKET_L type BRACKET_R #ListType
    | type BANG #NonNullType
    | value=IDEN #NamedType;

selectionSet : BRACE_L selections+=selection+ BRACE_R;

selection : STRING # selectString
    | field # selectField;

field: (alias=IDEN COLON)? name=IDEN argumentList? selectionSet?;

argumentList : PAREN_L arguments+=argument (COMMA arguments+=argument)* PAREN_R;

argument : name=IDEN COLON DOLLAR value=IDEN #argumentVariable
    | name=IDEN COLON value=STRING #argumentString;
