///
// File:  token.go
// Author: ymiyamoto
//
// Created on Wed Jun 20 00:17:56 2018
//
package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	SEMICOLON = ";"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywards = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywards[ident]; ok {
		return tok
	}
	return IDENT
}
