package repl

/*
#cgo CFLAGS: -I../flex/
#cgo LDFLAGS: -L../flex/ -lfl

#include <stdlib.h>
#include <stdio.h>
#include "../flex/lex.yy.c"

extern int yylex();
extern void yyset_in(FILE *input_file);
extern char* yyget_text();
*/
import "C"

import (
	"log"
	"os"

	"github.com/ZeroBl21/go-monkey-visualizer/internal/token"
)

// ParseTokensFlex es la función que ejecuta el lexer Flex y devuelve una lista de tokens
func ParseTokensFlex(input string) ([]token.Token, error) {
	// Crear un archivo temporal para que Flex lo lea
	tmpFile, err := os.CreateTemp("", "input-*.txt")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(input)
	if err != nil {
		return nil, err
	}
	tmpFile.Seek(0, 0)

	// Abrir el archivo con Flex
	C.yyset_in(C.fdopen(C.int(tmpFile.Fd()), C.CString("r")))

	var tokens []token.Token
	tokenMap := map[int]string{
		1:  "ILLEGAL",
		2:  "EOF",
		3:  "IDENT",
		4:  "INT",
		5:  "STRING",
		6:  "ASSIGN",
		7:  "PLUS",
		8:  "MINUS",
		9:  "BANG",
		10: "ASTERISK",
		11: "SLASH",
		12: "EQ",
		13: "NOT_EQ",
		14: "LT",
		15: "RT",
		16: "COMMA",
		17: "COLON",
		18: "SEMICOLON",
		19: "LPAREN",
		20: "RPAREN",
		21: "LBRACE",
		22: "RBRACE",
		23: "LBRACKET",
		24: "RBRACKET",
		25: "FUNCTION",
		26: "LET",
		27: "TRUE",
		28: "FALSE",
		29: "IF",
		30: "ELSE",
		31: "RETURN",
	}

	count := 0
	// Invocar yylex repetidamente
	for {
		tokenType := int(C.yylex())
		if tokenType == 0 { // EOF
			break
		}

		// Obtener el texto del token desde Flex
		literal := C.GoString(C.yyget_text())

		log.Printf("Type: %s; Literal: %s; Count: %d; TokenTypeNumber: %d; TokenType: %s;",
			tokenMap[tokenType],
			literal,
			count,
			tokenType,
			token.TokenType(tokenMap[tokenType]),
		)
		count++

		// Añadir el token a la lista de tokens
		tokens = append(tokens, token.Token{
			Type:    token.TokenType(tokenMap[tokenType]),
			Literal: literal,
		})
	}

	return tokens, nil
}
