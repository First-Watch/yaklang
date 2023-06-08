parser grammar SuricataRuleParser;

@header {
#pragma warning disable 3021
#pragma warning disable 0108
}

options { tokenVocab=SuricataRuleLexer; }

rules : rule+ EOF;
rule : action protocol src_address src_port direction dest_address dest_port params;

action : ID;
protocol : ID;
direction : (DirectionForward | DirectionBoth);

/* parse address */
src_address: address;
dest_address : address;
address
    : Any
    | Negative address
    | environment_var
    | ipv4
    | ipv6
    | LBracket address (Comma address)* RBracket
    ;
environment_var: Dollar ID;

/* ipv4 */
ipv4: ipv4block Dot ipv4block Dot ipv4block Dot ipv4block ('/' ipv4mask )? ;
ipv4block: INT;
ipv4mask: INT;

/* ipv6 */
ipv6 : hex_part (Colon hex_part)* (DoubleColon (hex_part Colon)* hex_part)?;
hex_part : h16 | h16? DoubleColon h16 | DoubleColon h16;
h16 : HEX;

/* ports */
src_port : port;
dest_port : port;
port
    : Any
    | INT
    | INT Colon INT ?
    | Colon INT
    | Negative port
    | LBracket port (Comma port)* RBracket
    | environment_var
    ;

/* rules configuration */
params : ParamStart param (ParamSep param)* ParamSep? ParamEnd;
param: (ParamKey ParamKeySep ParamValue) | (ParamKey) ;