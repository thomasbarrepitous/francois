package francois

import "strings"

type Token struct {
	kind  TokenKind
	value string
}

type TokenKind int

const (
	EOF TokenKind = iota
	Whitespace
	Number
	Identifier
	Operator
	Assignment
	OpenParen
	CloseParen
)

func identifyToken(word string) TokenKind {
	switch word {
	case " ":
		return Whitespace
	case "\n":
		return Operator
	case "+", "-", "*", "/", "%":
		return Assignment
	case "(":
		return OpenParen
	case ")":
		return CloseParen
	default:
		return Identifier
	}
}

func Tokenize(sourceCode string) (tokens []Token) {
	sourceCodeSplit := strings.Split(sourceCode, " ")
	for _, token := range sourceCodeSplit {
		tokens = append(tokens, Token{kind: Identifier, value: token})
	}
	return tokens
}
