package lexer

import (
	"strconv"
	"strings"
)

type Token struct {
	Kind  TokenKind
	Value string
}

type TokenKind string

const (
	EOF                 TokenKind = "EOF"
	Whitespace          TokenKind = "Whitespace"
	EndOfInstruction    TokenKind = "EndOfInstruction"
	Number              TokenKind = "Number"
	Operator            TokenKind = "Operator"
	OpenParen           TokenKind = "OpenParen"
	CloseParen          TokenKind = "CloseParen"
	Undefined           TokenKind = "Undefined"
	Identifier          TokenKind = "Identifier"
	Assignment          TokenKind = "Assignment"
	LocalDeclaration    TokenKind = "LocalDeclaration"
	ConstantDeclaration TokenKind = "ConstantDeclaration"
	Null                TokenKind = "Null"
)

var Keywords = map[string]TokenKind{
	"met":   LocalDeclaration,
	"const": ConstantDeclaration,
	"dans":  Assignment,
	"null":  Null,
	"fin":   EndOfInstruction,
}

func shift(sourceCode *string) string {
	if len(*sourceCode) > 0 {
		current := (*sourceCode)[0]
		*sourceCode = (*sourceCode)[1:]
		return string(current)
	}
	return ""
}

func Tokenize(sourceCode string) (tokens []Token) {
	for len(sourceCode) > 0 {
		token := string(sourceCode[0])
		switch token {
		case " ", "\t", "\n":
			shift(&sourceCode)
		case "+", "-", "*", "/", "%":
			tokens = append(tokens, Token{Kind: Operator, Value: shift(&sourceCode)})
		case "(":
			tokens = append(tokens, Token{Kind: OpenParen, Value: shift(&sourceCode)})
		case ")":
			tokens = append(tokens, Token{Kind: CloseParen, Value: shift(&sourceCode)})
		default:
			// Check for multi-character tokens
			word := extractWord(&sourceCode)
			tokens = append(tokens, Token{Kind: identifyMultiCharacterToken(word), Value: word})
		}
	}
	tokens = append(tokens, Token{Kind: EOF, Value: ""})
	return tokens
}

func identifyMultiCharacterToken(token string) TokenKind {
	if isAlpha(token) {
		if kind, ok := Keywords[token]; ok {
			return kind
		}
		return Identifier
	}
	if isNumber(token) {
		return Number
	}
	return Undefined
}

func extractWord(sourceCode *string) string {
	word := string(shift(sourceCode))
	for len(*sourceCode) > 0 {
		nextChar := string((*sourceCode)[0])
		if isAlpha(nextChar) || isNumber(nextChar) {
			word += shift(sourceCode)
		} else {
			break
		}
	}
	return word
}

func isAlpha(char string) bool {
	return strings.ToLower(char) != strings.ToUpper(char)
}

func isNumber(char string) bool {
	_, err := strconv.ParseFloat(char, 64)
	return err == nil
}
