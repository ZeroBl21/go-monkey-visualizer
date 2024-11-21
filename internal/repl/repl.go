package repl

import (
	"fmt"

	"github.com/ZeroBl21/go-monkey-visualizer/internal/ast"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/compiler"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/evaluator"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/lexer"
	"github.com/ZeroBl21/go-monkey-visualizer/internal/object"
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

type REPL struct {
	env *object.Environment
}

func New() *REPL {
	return &REPL{
		env: object.NewEnvironment(),
	}
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
	Program  *ast.Program `json:"program"`
	Errors   []string     `json:"errors"`
	Evaluate string       `json:"evaluate"`
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

func (r *REPL) EvaluateLine(line string) *ParseResult {
	l := lexer.New(line)
	p := parser.New(l)

	program := p.ParseProgram()
	result := &ParseResult{
		Program: program,
		Errors:  p.Errors(),
	}

	if len(result.Errors) != 0 {
		return result
	}

	evaluated := evaluator.Eval(program, r.env)
	if evaluated != nil {
		result.Evaluate = evaluated.Inspect()
	}

	return result
}

func (r *REPL) CompileToBytecode(line string) (*compiler.Bytecode, error) {
	l := lexer.New(line)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return nil, fmt.Errorf("parser errors: %v", p.Errors())
	}

	comp := compiler.New()
	if err := comp.Compile(program); err != nil {
		return nil, fmt.Errorf("compiler error: %s", err)
	}

	return comp.Bytecode(), nil
}

func applyColor(color, text string) string {
	return color + text + RESET
}
