lexer grammar SuricataRuleLexer;

@header {
#pragma warning disable 3021
}

// Keywords
Any: 'any';

// Symbols
Negative: '!';
Dollar: '$';
DirectionForward: '->';
DirectionBoth: '<>';
Mul: '*';
Div: '/';
Mod: '%';
Amp: '&';
Plus: '+';
Sub: '-';
Power: '^';
Lt: '<';
Gt: '>';
LtEq: '<=';
GtEq: '>=';
SemiColon: ';';
Colon: ':';
DoubleColon: '::';
LBracket: '[';
RBracket: ']';
ParamStart: '(' -> pushMode(PARAM_MODE);
LBrace: '{';
RBrace: '}';
Comma: ',';
Eq: '=';
NotSymbol: '~';
Dot: '.';

LINE_COMMENT: ('#' | '//') SingleLineInputCharacter* -> skip;

ID
    : [a-zA-Z_][a-zA-Z_0-9]*
    ;

NORMALSTRING
    : '"' ( EscapeSequence | ~('"') )* '"'
    ;

INT
    : Digit+
    ;

HEX
    : HexDigit+
    ;

fragment
ExponentPart
    : [eE] [+-]? Digit+
    ;

fragment
HexExponentPart
    : [pP] [+-]? Digit+
    ;

fragment
EscapeSequence
    : '\\' [abfnrtvz"'|$#\\]   // World of Warcraft Lua additionally escapes |$#
    | '\\' '\r'? '\n'
    | DecimalEscape
    | HexEscape
    | UtfEscape
    ;
fragment
DecimalEscape
    : '\\' Digit
    | '\\' Digit Digit
    | '\\' [0-2] Digit Digit
    ;
fragment
HexEscape
    : '\\' 'x' HexDigit HexDigit
    ;
fragment
UtfEscape
    : '\\' 'u{' HexDigit+ '}'
    ;
fragment
Digit
    : [0-9]
    ;
fragment
HexDigit
    : [0-9a-fA-F]
    ;
fragment
SingleLineInputCharacter
    : ~[\r\n\u0085\u2028\u2029]
    ;
WS
    : [ \t\u000C\r\n]+ -> skip
    ;
NonSemiColon: [^;]+;
SHEBANG
    : '#' '!' SingleLineInputCharacter* -> channel(HIDDEN)
    ;
    
mode PARAM_VALUE;
    ParamNot: '!';
    ParamValue: ParamNot? ParamKey -> popMode;
    
mode PARAM_MODE;
    fragment Quote: '"';
    ParamKey: NORMALSTRING | (FreeValueAnyCharNoColon+) ;
    ParamKeySep: (':') -> pushMode(PARAM_VALUE);
    ParamSep: ';';
    ParamEnd: ')' -> popMode;
    fragment FreeValueAnyCharNoColon: ~(')' | '"' | ';' | ':');