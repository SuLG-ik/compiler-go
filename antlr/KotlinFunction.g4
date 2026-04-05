grammar KotlinFunction;

program
    : funDecl SEMI EOF
    ;

funDecl
    : FUN IDENTIFIER LPAREN paramList? RPAREN COLON typeRef LBRACE body RBRACE
    ;

paramList
    : param (COMMA param)*
    ;

param
    : IDENTIFIER COLON typeRef
    ;

typeRef
    : INT_TYPE
    | BOOLEAN_TYPE
    | FLOAT_TYPE
    | DOUBLE_TYPE
    ;

body
    : returnStmt
    ;

returnStmt
    : RETURN expr
    ;

expr
    : term ((PLUS | MINUS) term)*
    ;

term
    : factor ((STAR | SLASH) factor)*
    ;

factor
    : IDENTIFIER
    | FLOAT_LITERAL
    | INTEGER_LITERAL
    | LPAREN expr RPAREN
    ;

FUN: 'fun';
RETURN: 'return';
INT_TYPE: 'Int';
BOOLEAN_TYPE: 'Boolean';
FLOAT_TYPE: 'Float';
DOUBLE_TYPE: 'Double';
COLON: ':';
COMMA: ',';
SEMI: ';';
PLUS: '+';
MINUS: '-';
STAR: '*';
SLASH: '/';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
FLOAT_LITERAL: [0-9]+ '.' [0-9]+;
INTEGER_LITERAL: [0-9]+;
IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]*;
WS: [ \t\r\n]+ -> skip;