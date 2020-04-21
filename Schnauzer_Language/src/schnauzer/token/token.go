package token

type TokenType string

const(
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  //Identifiers + literals
  IDENT = "IDENT" //add, x, y
  INT = "INT" // 214234135324

  //Operators
  ASSIGN = "="
  PLUS = "+"

  //Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"

)

type Token struct {
  
  Type  TokenType
  Literal string

}
