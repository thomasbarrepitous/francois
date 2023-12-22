package parser

import (
	"fmt"
	"francois/ast"
	"francois/lexer"
	"log"
	"os"
	"strconv"
)

type Parser struct {
	tokens []lexer.Token
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
		Type: ast.ProgramType,
		Body: []ast.Statement{},
	}
	fmt.Printf("%+v\n", p.tokens)

	for !p.isEOF() {
		program.Body = append(program.Body, p.parseStatement())
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	return p.parseExpression()
}

func (p *Parser) parseExpression() ast.Expression {
	return p.parseAdditiveExpression()
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	switch p.peek().Kind {
	case lexer.Identifier:
		return ast.Identifier{
			Type:   ast.IdentifierType,
			Symbol: p.consume().Value,
		}
	case lexer.Number:
		value, err := strconv.ParseFloat(p.consume().Value, 64)
		if err != nil {
			log.Fatal("Error parsing float!", err)
			os.Exit(1)
		}
		return ast.NumericLiteral{
			Type:  ast.NumericLiteralType,
			Value: value,
		}
	default:
		log.Fatal("Unexpected token found during parsing!", p.peek())
		os.Exit(1)
		return nil
	}
}

func (p *Parser) parseAdditiveExpression() ast.Expression {
	left := p.parsePrimaryExpression()
	for p.peek().Value == "+" || p.peek().Value == "-" {
		operator := p.consume()
		right := p.parsePrimaryExpression()
		left = ast.BinaryExpression{
			Type:     ast.BinaryExpressionType,
			Operator: operator.Value,
			Left:     left,
			Right:    right,
		}
	}
	return left
}
