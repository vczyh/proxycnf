grammar ProxySQL;

configuration
    : settingList? EOF
    ;

settingList
    : setting
    | settingList setting
    ;

setting
    : name (':' | '=') value (';' | ',' |)
    ;

name: ID;

value
    : scalarValue
    | array
    | list
    | group
    ;

valueList
    : value
    | valueList (',' | ) value
    ;

// todo more
scalarValue
    : BOOLEAN
    | NUMBER
    | FLOAT
    | HEX
    | STRING
    ;

scalarValueList
    : scalarValue
    | scalarValueList ',' scalarValue
    ;

array
    : '[' scalarValueList? ']'
    ;

list
    : '(' valueList? ')'
    ;

group
    : '{' settingList? '}'
    ;


BOOLEAN: 'true' | 'false' | 'TRUE' | 'FALSE';
ID: [a-zA-Z_][a-zA-Z0-9_]*;
STRING: '"' (~["\\] | '\\' .)* '"';
NUMBER: [0-9]+;
FLOAT: [0-9]+ '.' [0-9]* ([eE] [+-]? [0-9]+)?;
//BOOLEAN: 'true' | 'false';
HEX: '0x' [0-9a-fA-F]+;

COMMENT: '#' ~[\r\n]* -> skip;
WS: [ \t\r\n]+ -> skip;

