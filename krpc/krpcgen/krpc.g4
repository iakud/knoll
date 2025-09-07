// Define a grammar called krpc
grammar krpc;

krpc
	: packageStatement (topLevelDef | emptyStatement_)* EOF
	;

// Package

packageStatement
	: PACKAGE fullIdent SEMI
	;

// Normal Field

field
	: fieldLabel? type_ fieldName EQ fieldNumber (LB fieldOptions RB)? SEMI
	;

fieldLabel
	: REPEATED
	;

fieldOptions
	: fieldOption (COMMA fieldOption)*
	;

fieldOption
	: fullIdent
	;

fieldNumber
	: intLit
	;

// Map field

mapField
	: MAP LT keyType COMMA type_ GT mapName EQ fieldNumber (LB fieldOptions RB)? SEMI
	;

keyType
	: INT32
	| INT64
	| UINT32
	| UINT64
	| SINT32
	| SINT64
	| FIXED32
	| FIXED64
	| SFIXED32
	| SFIXED64
	| BOOL
	| STRING
	;

// field types

type_
	: DOUBLE
	| FLOAT
	| INT32
	| INT64
	| UINT32
	| UINT64
	| SINT32
	| SINT64
	| FIXED32
	| FIXED64
	| SFIXED32
	| SFIXED64
	| BOOL
	| STRING
	| BYTES
	| TIMESTAMP
	| DURATION
	| EMPTY
	| messageType
	| enumType
	;

// Top Level definitions

topLevelDef
	: rpcDef
	| ntfDef
	;

// rpc

rpcDef
	: RPC rpcName LP rpcRequestBody RP EQ intLit
	;

// ntf

ntfDef
	: NTF ntfName LP ntfBody RP EQ intLit
	;

// lexical

emptyStatement_
	: SEMI
	;

// Lexical elements

ident
	: IDENTIFIER
	| keywords
	;

rpcName
	: ident
	;

ntfName
	: ident
	;

intLit
	: INT_LIT
	;

// keywords
RPC
	: 'rpc'
	;

NTF
	: 'ntf'
	;

// symbols

SEMI
	: ';'
	;

EQ
	: '='
	;

LP
	: '('
	;

RP
	: ')'
	;

INT_LIT
	: DECIMAL_LIT
	| OCTAL_LIT
	| HEX_LIT
	;

fragment DECIMAL_LIT
	: ([1-9]) DECIMAL_DIGIT*
	;

fragment OCTAL_LIT
	: '0' OCTAL_DIGIT*
	;

fragment HEX_LIT
	: '0' ('x' | 'X') HEX_DIGIT+
	;

IDENTIFIER
	: LETTER (LETTER | DECIMAL_DIGIT)*
	;

fragment LETTER
	: [A-Za-z_]
	;

fragment DECIMAL_DIGIT
	: [0-9]
	;

fragment OCTAL_DIGIT
	: [0-7]
	;

fragment HEX_DIGIT
	: [0-9A-Fa-f]
	;

// comments
WS
	: [ \t\r\n\u000C]+ -> skip
	;

LINE_COMMENT
	: '//' ~[\r\n]* -> channel(HIDDEN)
	;

COMMENT
	: '/*' .*? '*/' -> channel(HIDDEN)
	;

keywords
	: SYNTAX
	| IMPORT
	| PROTO_GO_PACKAGE
	| WEAK
	| PUBLIC
	| PACKAGE
	| OPTION
	| OPTIONAL
	| REPEATED
	| ONEOF
	| MAP
	| INT32
	| INT64
	| UINT32
	| UINT64
	| SINT32
	| SINT64
	| FIXED32
	| FIXED64
	| SFIXED32
	| SFIXED64
	| BOOL
	| STRING
	| DOUBLE
	| FLOAT
	| BYTES
	| TIMESTAMP
	| DURATION
	| EMPTY
	| RESERVED
	| TO
	| MAX
	| ENUM
	| ENTITY
	| COMPONENT
	| MESSAGE
	| SERVICE
	| EXTEND
	| RPC
	| STREAM
	| RETURNS
	| BOOL_LIT
	;