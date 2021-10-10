lexer grammar QuantyLexer;
tokens { STRING }

// KEYWORDS
COMPONENT : 'component';

// STRINGS
DQUOTE : '"' .*? '"' -> type(STRING);

// COMMENTS
BLOCK_COMMENT
	: '/*' .*? '*/' -> channel(HIDDEN)
	;
LINE_COMMENT
	: '//' ~[\r\n]* -> channel(HIDDEN)
	;

// BASE
LBRACE : '{';
RBRACE : '}';
LPAREN : '(';
RPAREN : ')';
COMMA : ',';
COLON : ':';

IDEN : [a-zA-Z][a-zA-Z0-9]*;

INT : DIGIT+;

fragment DIGIT: [0-9];

WS : [ \r\n\t]+ -> channel(HIDDEN);
