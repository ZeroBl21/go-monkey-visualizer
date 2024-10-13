%{
#include <stdio.h>
#include <stdlib.h>

void yyerror(const char *s);

#define ILLEGAL     1
#define EOF_TOK     2
#define IDENT       3
#define INT         4
#define STRING      5

#define ASSIGN      6
#define PLUS        7
#define MINUS       8
#define BANG        9
#define ASTERISK    10
#define SLASH       11

#define EQ          12
#define NOT_EQ      13
#define LT          14
#define RT          15

#define COMMA       16
#define COLON       17
#define SEMICOLON   18
#define LPAREN      19
#define RPAREN      20
#define LBRACE      21
#define RBRACE      22
#define LBRACKET    23
#define RBRACKET    24

#define FUNCTION    25
#define LET         26
#define TRUE        27
#define FALSE       28
#define IF          29
#define ELSE        30
#define RETURN      31
%}

%option noyywrap

DIGIT   [0-9]
LETTER  [a-zA-Z_]
OP      [\+\-\*/=<>!]

%%

"fn"          { return FUNCTION; }
"let"         { return LET; }
"true"        { return TRUE; }
"false"       { return FALSE; }
"if"          { return IF; }
"else"        { return ELSE; }
"return"      { return RETURN; }

{LETTER}({LETTER}|{DIGIT})*   { return IDENT; }
{DIGIT}+                      { return INT; }

\"(\\.|[^\"])*\"              { return STRING; }

"="           { return ASSIGN; }
"=="          { return EQ; }
"!="          { return NOT_EQ; }
"<"           { return LT; }
">"           { return RT; }
"+"           { return PLUS; }
"-"           { return MINUS; }
"*"           { return ASTERISK; }
"/"           { return SLASH; }
"!"           { return BANG; }

","           { return COMMA; }
":"           { return COLON; }
";"           { return SEMICOLON; }
"("           { return LPAREN; }
")"           { return RPAREN; }
"{"           { return LBRACE; }
"}"           { return RBRACE; }
"["           { return LBRACKET; }
"]"           { return RBRACKET; }

[ \t\n\r]+    { /* ignore whitespace */ }

"//".*        { /* skip comments */ }

.             { return ILLEGAL; }

%%

void yyerror(const char *s) {
    fprintf(stderr, "Error: %s\n", s);
}
