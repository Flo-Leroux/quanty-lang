lexer grammar Lexer;

tokens { STRING }

// KEYWORDS
MODULE : 'module';
COMPONENT : 'component';

// STRINGS
DQUOTE : '"' .*? '"' -> type(STRING);

// COMMENTS
COMMENT : '#' ~[\r\n]* -> channel(HIDDEN);

// BASE
DOLLAR : '$';
LBRACE : '{';
RBRACE : '}';
LPAREN : '(';
RPAREN : ')';
COMMA : ',';
COLON : ':';
DOT : '.';

IDEN : [a-zA-Z][a-zA-Z0-9_]*;

ID : '"' [a-zA-Z][a-zA-Z0-9_-]* '"';
INT : DIGIT+;
FLOAT : DIGIT+('.'|',')DIGIT+;
BOOLEAN : 'true' | 'false';

fragment DIGIT : [0-9];

WS : [ \r\n\t]+ -> channel(HIDDEN);
