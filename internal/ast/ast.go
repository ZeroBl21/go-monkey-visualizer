package ast

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/ZeroBl21/go-monkey-visualizer/internal/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (s *LetStatement) statementNode()       {}
func (s *LetStatement) TokenLiteral() string { return s.Token.Literal }
func (s *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Name.String())
	out.WriteString(" = ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (s *ReturnStatement) statementNode()       {}
func (s *ReturnStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")

	if s.ReturnValue != nil {
		out.WriteString(s.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (s *ExpressionStatement) statementNode()       {}
func (s *ExpressionStatement) TokenLiteral() string { return s.Token.Literal }
func (s *ExpressionStatement) String() string {
	if s.Expression != nil {
		return s.Expression.String()
	}

	return ""
}

type BlockStatement struct {
	Token      token.Token // The { Token
	Statements []Statement
}

func (s *BlockStatement) statementNode()       {}
func (s *BlockStatement) TokenLiteral() string { return s.Token.Literal }
func (s *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range s.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Literals

type Identifier struct {
	Token token.Token // The 'Return' token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (l *IntegerLiteral) expressionNode()      {}
func (l *IntegerLiteral) TokenLiteral() string { return l.Token.Literal }
func (l *IntegerLiteral) String() string       { return l.Token.Literal }

type StringLiteral struct {
	Token token.Token
	Value string
}

func (l *StringLiteral) expressionNode()      {}
func (l *StringLiteral) TokenLiteral() string { return l.Token.Literal }
func (l *StringLiteral) String() string       { return l.Token.Literal }

type Boolean struct {
	Token token.Token
	Value bool
}

func (l *Boolean) expressionNode()      {}
func (l *Boolean) TokenLiteral() string { return l.Token.Literal }
func (l *Boolean) String() string       { return l.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token // The '[' token
	Elements []Expression
}

func (l *ArrayLiteral) expressionNode()      {}
func (l *ArrayLiteral) TokenLiteral() string { return l.Token.Literal }
func (l *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, p := range l.Elements {
		elements = append(elements, p.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type HashPairs map[Expression]Expression

func (hp HashPairs) MarshalJSON() ([]byte, error) {
	type Pair struct {
		Key   string
		Value Expression
	}
	var pairs []Pair

	for key, value := range hp {
		pairs = append(pairs, Pair{
			Key:   key.String(),
			Value: value,
		})
	}

	return json.Marshal(pairs)
}

type HashLiteral struct {
	Token token.Token // The '{' token
	Pairs HashPairs
}

func (l *HashLiteral) expressionNode()      {}
func (l *HashLiteral) TokenLiteral() string { return l.Token.Literal }
func (l *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range l.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (l *FunctionLiteral) expressionNode()      {}
func (l *FunctionLiteral) TokenLiteral() string { return l.Token.Literal }
func (l *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range l.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(l.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(l.Body.String())

	return out.String()
}

// Expressions

type IndexExpression struct {
	Token token.Token // The '[' Token
	Left  Expression
	Index Expression
}

func (e *IndexExpression) expressionNode()      {}
func (e *IndexExpression) TokenLiteral() string { return e.Token.Literal }
func (e *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString("[")
	out.WriteString(e.Index.String())
	out.WriteString("])")

	return out.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  //  Identifier or FunctionLiteral
	Arguments []Expression
}

func (e *CallExpression) expressionNode()      {}
func (e *CallExpression) TokenLiteral() string { return e.Token.Literal }
func (e *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range e.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(e.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (e *PrefixExpression) expressionNode()      {}
func (e *PrefixExpression) TokenLiteral() string { return e.Token.Literal }
func (e *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(e.Operator)
	out.WriteString(e.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // Operator token, e.g +
	Left     Expression
	Operator string
	Right    Expression
}

func (e *InfixExpression) expressionNode()      {}
func (e *InfixExpression) TokenLiteral() string { return e.Token.Literal }
func (e *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(e.Left.String())
	out.WriteString(" " + e.Operator + " ")
	out.WriteString(e.Right.String())
	out.WriteString(")")

	return out.String()
}

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (e *IfExpression) expressionNode()      {}
func (e *IfExpression) TokenLiteral() string { return e.Token.Literal }
func (e *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(e.Condition.String())
	out.WriteString(" ")
	out.WriteString(e.Consequence.String())

	if e.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(e.Alternative.String())
	}

	return out.String()
}
