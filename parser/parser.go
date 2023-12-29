package parser

import (
	"francois/ast"
	"francois/lexer"
	"log"
	"os"
	"strconv"
)

type Parser struct {
	tokens []lexer.Token
}

func (p *Parser) MustConsume(kind lexer.TokenKind, message string) lexer.Token {
	if p.peek().Kind != kind {
		log.Fatal(message)
		os.Exit(1)
	}
	return p.consume()
}

func (p *Parser) isEOF() bool {
	return p.tokens[0].Kind == lexer.EOF
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[0]
}

func (p *Parser) consume() lexer.Token {
	token := p.tokens[0]
	p.tokens = p.tokens[1:]
	return token
}

func (p *Parser) ProduceAST(sourceCode string) *ast.Program {
	p.tokens = lexer.Tokenize(sourceCode)
	program := &ast.Program{
		Body: []ast.Statement{},
	}

	for !p.isEOF() {
		program.Body = append(program.Body, p.parseStatement())
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.peek().Kind {
	case lexer.LocalDeclaration:
		return p.parseVariableDeclaration()
	default:
		return p.parseExpression()
	}
}

func (p *Parser) parseExpression() ast.Expression {
	return p.parseAssignmentExpression()
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	switch p.peek().Kind {
	case lexer.Identifier:
		return &ast.Identifier{
			Symbol: p.consume().Value,
		}
	case lexer.Null:
		p.consume()
		return &ast.NullLiteral{}
	case lexer.Number:
		value, err := strconv.ParseFloat(p.consume().Value, 64)
		if err != nil {
			log.Fatal("Error parsing float!", err)
			os.Exit(1)
		}
		return &ast.NumericLiteral{
			Value: value,
		}
	case lexer.OpenParen:
		p.consume() // opening paren
		value := p.parseExpression()
		p.MustConsume(
			lexer.CloseParen,
			"Unexpected token. Expected closing parenthesis.",
		)
		return value
	default:
		log.Fatal("Unexpected token found during parsing!", p.peek())
		os.Exit(1)
		return nil
	}
}

func (p *Parser) parseVariableDeclaration() ast.Statement {
	kind := p.consume().Kind
	identifier := p.MustConsume(
		lexer.Identifier,
		"Unexpected token. Expected identifier.",
	)
	// Declaration without default value.
	if p.peek().Kind == lexer.EndOfInstruction {
		if kind == lexer.ConstantDeclaration {
			log.Fatal("Unexpected token. Expected assignment.")
			os.Exit(1)
		}
		return &ast.VariableDeclaration{
			IsConstant: false,
			Identifier: ast.Identifier{Symbol: identifier.Value},
			Value:      &ast.NullLiteral{},
		}
	}
	p.MustConsume(
		lexer.Assignment,
		"Unexpected token. Expected assignment.",
	)
	value := p.parseExpression()
	declaration := &ast.VariableDeclaration{
		IsConstant: kind == lexer.ConstantDeclaration,
		Identifier: ast.Identifier{Symbol: identifier.Value},
		Value:      value,
	}
	p.MustConsume(
		lexer.EndOfInstruction,
		"Unexpected token. Expected end of instruction.",
	)
	return declaration
}

func (p *Parser) parseMultiplicativeExpression() ast.Expression {
	left := p.parsePrimaryExpression()
	for p.peek().Value == "/" || p.peek().Value == "*" || p.peek().Value == "%" {
		operator := p.consume()
		right := p.parsePrimaryExpression()
		left = &ast.BinaryExpression{
			Operator: operator.Value,
			Left:     left,
			Right:    right,
		}
	}
	return left
}

func (p *Parser) parseAdditiveExpression() ast.Expression {
	left := p.parseMultiplicativeExpression()
	for p.peek().Value == "+" || p.peek().Value == "-" {
		operator := p.consume()
		right := p.parseMultiplicativeExpression()
		left = &ast.BinaryExpression{
			Operator: operator.Value,
			Left:     left,
			Right:    right,
		}
	}
	return left
}

func (p *Parser) parseAssignmentExpression() ast.Expression {
	left := p.parseAdditiveExpression()
	if p.peek().Kind == lexer.Assignment {
		// Get rid of the assignment token.
		p.consume()
		value := p.parseAssignmentExpression()
		return &ast.AssignmentExpression{
			Assignee: left,
			Value:    value,
		}
	}
	return left
}
