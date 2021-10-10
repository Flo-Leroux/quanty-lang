grammar QuantyParser;

prog : component | EOF;

component : 'component' IDEN selectionset;

selectionset : '{' tag* '}';

tag : IDEN (selectionset)?;

IDEN : [a-zA-Z][a-zA-Z0-9]*;

WS : [ \r\n\t]+ -> skip;
