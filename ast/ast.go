package ast

import (
	"francois/runtime"
	"log"
)

// TokenType represents the type of AST token.
type TokenType string

const (
	ProgramToken             TokenType = "Program"
	NumericLiteralToken      TokenType = "NumericLiteral"
	NullLiteralToken         TokenType = "NullLiteral"
	VariableDeclarationToken TokenType = "VariableDeclaration"
	IdentifierToken          TokenType = "Identifier"
	BinaryExpressionToken    TokenType = "BinaryExpr"
)

// Statement represents a statement in the AST.
type Statement interface {
	Kind() TokenType
	Evaluate(*runtime.Environment) runtime.RuntimeValue
}

// Expression represents an expression in the AST.
type Expression interface {
	Statement
}

// Program represents a block of statements.
type Program struct {
	Body []Statement
}

func (p *Program) Kind() TokenType {
	return ProgramToken
}

func (program *Program) Evaluate(env *runtime.Environment) runtime.RuntimeValue {
	var lastEvaluated runtime.RuntimeValue
	for _, statement := range program.Body {
		lastEvaluated = statement.Evaluate(env)
	}
	return lastEvaluated
}

// BinaryExpression represents a binary expression with an operator between two expressions.
type BinaryExpression struct {
	Expression
	Left     Expression
	Right    Expression
	Operator string
}

func (b *BinaryExpression) Kind() TokenType {
	return BinaryExpressionToken
}

func (binOp *BinaryExpression) Evaluate(env *runtime.Environment) runtime.RuntimeValue {
	left := binOp.Left.Evaluate(env)
	right := binOp.Right.Evaluate(env)
	if left.Type() == runtime.NumericType && right.Type() == runtime.NumericType {
		switch binOp.Operator {
		case "+":
			return runtime.MakeNumericValue(left.(runtime.NumericValue).Value + right.(runtime.NumericValue).Value)
			// return runtime.NumericValue{Value: left.(runtime.NumericValue).Value + right.(runtime.NumericValue).Value}
		case "-":
			return runtime.MakeNumericValue(left.(runtime.NumericValue).Value - right.(runtime.NumericValue).Value)
			// return runtime.NumericValue{Value: left.(runtime.NumericValue).Value - right.(runtime.NumericValue).Value}
		case "*":
			return runtime.MakeNumericValue(left.(runtime.NumericValue).Value * right.(runtime.NumericValue).Value)
			// return runtime.NumericValue{Value: left.(runtime.NumericValue).Value * right.(runtime.NumericValue).Value}
		case "/":
			return runtime.MakeNumericValue(left.(runtime.NumericValue).Value / right.(runtime.NumericValue).Value)
			// return runtime.NumericValue{Value: left.(runtime.NumericValue).Value / right.(runtime.NumericValue).Value}
		case "%":
			return runtime.MakeNumericValue(float64(int(left.(runtime.NumericValue).Value) % int(right.(runtime.NumericValue).Value)))
			// return runtime.NumericValue{Value: float64(int(left.(runtime.NumericValue).Value) % int(right.(runtime.NumericValue).Value))}
		default:
			log.Fatal("Unknown operator")
			return runtime.MakeNullValue()
		}
	}
	return runtime.MakeNullValue()
}

// Identifier represents a user-defined variable or symbol.
type Identifier struct {
	Symbol string
}

func (i *Identifier) Kind() TokenType {
	return IdentifierToken
}

func (identifier *Identifier) Evaluate(env *runtime.Environment) runtime.RuntimeValue {
	value := env.GetVariable(identifier.Symbol)
	return value
}

type VariableDeclaration struct {
	IsConstant bool
	Value      Expression
	Identifier
}

func (v *VariableDeclaration) Kind() TokenType {
	return VariableDeclarationToken
}

func (variableDeclaration *VariableDeclaration) Evaluate(env *runtime.Environment) runtime.RuntimeValue {
	/*
		const value = declaration.value
		? evaluate(declaration.value, env)
		: MK_NULL();

		return env.declareVar(declaration.identifier, value, declaration.constant);
	*/
	if variableDeclaration.Value == nil {
		return env.DeclareVariable(variableDeclaration.Symbol, runtime.MakeNullValue())
	}
	value := variableDeclaration.Value.Evaluate(env)
	return env.DeclareVariable(variableDeclaration.Symbol, value)
}

// NumericLiteral represents a numeric constant in the source code.
type NumericLiteral struct {
	Value float64
}

func (n *NumericLiteral) Kind() TokenType {
	return NumericLiteralToken
}

func (numericLiteral *NumericLiteral) Evaluate(env *runtime.Environment) runtime.RuntimeValue {
	return runtime.MakeNumericValue(numericLiteral.Value)
}

type NullLiteral struct{}

func (n *NullLiteral) Kind() TokenType {
	return NullLiteralToken
}

func (nullLiteral *NullLiteral) Evaluate(env *runtime.Environment) runtime.RuntimeValue {
	return runtime.MakeNullValue()
}
