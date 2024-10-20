package repl

import (
	"github.com/ZeroBl21/go-monkey-visualizer/internal/ast"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/lexer"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/parser"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/token"
)

const (
	RESET  = "\033[0m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PROMPT = ">> "
)

const (
	CompileFlag = 1 << iota
	LexerFlag
	PrecedenceFlag
)

type REPL struct{}

func New() *REPL {
	return &REPL{}
}

func (r *REPL) ParseTokens(line string) []token.Token {
	var tokens []token.Token
	l := lexer.New(line)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}

	return tokens
}

type ParseResult struct {
	Program *ast.Program `json:"program"`
	Errors  []string     `json:"errors"`
}

func (r *REPL) ParseAST(line string) *ParseResult {
	l := lexer.New(line)
	p := parser.New(l)

	program := p.ParseProgram()

	result := &ParseResult{
		Program: program,
		Errors:  p.Errors(),
	}

	return result
}

func applyColor(color, text string) string {
	return color + text + RESET
}
