lexer grammar QuantyLexer;

channels { WSCHANNEL, MYHIDDEN }

tokens { STRING }
DOUBLE : '"' .*? '"'   -> type(STRING) ;

// Keywords
COMPONENT : 'component';

// Tokens
LBRACE : '{';
RBRACE : '}';
D_QUOTE : '"';

IDENT
    : [a-zA-Z][a-zA-Z0-9]*
    ;

// DOUBLE

// STRING : D_QUOTE (ESC | SAFECODEPOINT)* D_QUOTE;
//    : D_QUOTE (ESC | SAFECODEPOINT)* D_QUOTE
//    ;


fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;


fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;


fragment HEX
   : [0-9a-fA-F]
   ;


fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;


NUMBER
   : '-'? INT ('.' [0-9] +)? EXP?
   ;


fragment INT
   : '0' | [1-9] [0-9]*
   ;

// no leading zeros

fragment EXP
   : [Ee] [+\-]? INT
   ;

// \- since - means "range" inside [...]

BLOCK_COMMENT
	: '/*' .*? '*/' -> channel(HIDDEN)
	;
LINE_COMMENT
	: '//' ~[\r\n]* -> channel(HIDDEN)
	;

WS
   : [ \t\n\r] +  -> channel(WSCHANNEL)
   ;
