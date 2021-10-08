lexer grammar Lexer;

IDEN : [a-zA-Z][a-zA-Z0-9]*;

WS : [ \r\n\t]+ -> skip;
