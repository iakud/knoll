package internal

type TokenType byte

const (
	TokenNone TokenType = iota
	TokenPropertyName
	TokenString
	TokenNumber
	TokenTrue
	TokenFalse
	TokenNull
	TokenStartObject
	TokenEndObject
	TokenStartArray
	TokenEndArray
	TokenComment
)
