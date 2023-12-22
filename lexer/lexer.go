package lexer

import (
	"strconv"
	"strings"
)

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
	Declaration
	Undefined
)

var Keywords = map[string]TokenKind{
	"met":  Declaration,
	"dans": Assignment,
}

func Tokenize(sourceCode string) (tokens []Token) {
	sourceCodeSplit := strings.Split(sourceCode, " ")
	for _, token := range sourceCodeSplit {
		tokens = append(tokens, Token{kind: identifyToken(token), value: token})
	}
	return tokens
}

func identifyToken(word string) TokenKind {
	switch word {
	case " ":
		return Whitespace
	case "\n":
		return EOF
	case "+", "-", "*", "/", "%":
		return Operator
	case "(":
		return OpenParen
	case ")":
		return CloseParen
	default:
		if isInt(word) {
			return Number
		}
		if isAlpha(word) {
			if Keywords[word] != 0 {
				return Keywords[word]
			}
			return Identifier
		}
		return Undefined
	}
}

func isAlpha(char string) bool {
	return strings.ToLower(char) != strings.ToUpper(char)
}

func isInt(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}
