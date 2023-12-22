package ast

type TokenType string

const (
	ProgramType          TokenType = "Program"
	NumericLiteralType   TokenType = "NumericLiteral"
	IdentifierType       TokenType = "Identifier"
	BinaryExpressionType TokenType = "BinaryExpression"
)

type Statement Expression

type Expression interface{}

type Program struct {
	Body []Statement
	Type TokenType
}

type NumericLiteral struct {
	Value float64
	Type  TokenType
}

type Identifier struct {
	Symbol string
	Type   TokenType
}

type BinaryExpression struct {
	Operator string
	Left     Expression
	Right    Expression
	Type     TokenType
}
