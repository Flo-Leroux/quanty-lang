parser grammar Component;

import Lexer;

component : 'component' IDEN selectionset;

selectionset : '{' tag* '}';

tag : IDEN (selectionset)?;
