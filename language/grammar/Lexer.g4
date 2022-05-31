lexer grammar Lexer;

tokens { STRING }

// KEYWORDS
MODULE : M O D U L E;
COMPONENT : C O M P O N E N T;
FROM : F R O M;
IMPORT : I M P O R T;

// STRINGS
DQUOTE : '"' .*? '"' -> type(STRING);

// COMMENTS
COMMENT : '#' ~[\r\n]* -> channel(HIDDEN);

// BASE
BANG : '!';
DOLLAR : '$';
AMP: '&';
PAREN_L : '(';
PAREN_R : ')';
BRACE_L : '{';
BRACE_R : '}';
BRACKET_L: '[';
BRACKET_R: ']';
SPREAD: '...';
COLON : ':';
EQUALS: '=';
AT: '@';
PIPE: '|';
COMMA : ',';
DOT : '.';

IDEN : [a-zA-Z][a-zA-Z0-9_]*;

ID : '"' [a-zA-Z][a-zA-Z0-9_-]* '"';
INT : DIGIT+;
FLOAT : DIGIT+('.'|',')DIGIT+;
BOOLEAN : T R U E | F A L S E;

WS : [ \r\n\t]+ -> channel(HIDDEN);

fragment DIGIT : [0-9];

fragment A : [aA]; // match either an 'a' or 'A'
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
