package ast

type NodeType int

const (
	ProgramType NodeType = iota
	NumericLiteralType
	IdentifierType
	BinaryExpressionType
)

type Statement struct {
	kind NodeType
}

type Program struct {
	body []Statement
	Statement
}

type NumericLiteral struct {
	value float64
	Statement
}

type Identifier struct {
	symbol string
	Statement
}

type BinaryExpression struct {
	operator string
	left     Expression
	right    Expression
	Statement
}

type Expression struct {
	Statement
}
