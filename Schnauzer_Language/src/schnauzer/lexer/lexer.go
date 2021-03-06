package lexer

import "schnauzer/token"

type Lexer struct {
  input string
  position int
  readPosition int
  ch byte
}


func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
}

func (l *Lexer) NextToken() token.Token {
  var tok token.Token

  l.skipWhitespace() //Shnauzer DGAF about whitespace and calls the function to prove it

  switch l.ch {
  case '=':
    tok = newToken(token.ASSIGN, l.ch)
  case ';':
    tok = newToken(token.SEMICOLON, l.ch)
  case '(':
    tok = newToken(token.LPAREN, l.ch)
  case ')':
    tok = newToken(token.RPAREN, l.ch)
  case ',':
    tok = newToken(token.COMMA, l.ch)
  case '+':
    tok = newToken(token.PLUS, l.ch)
  case '{':
    tok = newToken(token.LBRACE, l.ch)
  case '}':
    tok = newToken(token.RBRACE, l.ch)
  case 0:
    tok.Literal = ""
    tok.Type = token.EOF
  default:
    if isLetter(l.ch) {
      tok.Literal = l.readIdentifier()
      tok.Type = token.LookupIdent(tok.Literal)  //Now goes to token package to lookup
      return tok
    } else if isDigit(l.ch) {
      tok.Type = token.INT
      tok.Literal = l.readNumber()
      return tok
      } else {
      tok = newToken(token.ILLEGAL, l.ch)
    }
  }
  l.readChar()
  return tok
}

//Schnauzer will now be able to read the Digits
func(l *Lexer) readNumber() string {
  position := l.position
  for isDigit(l.ch) {
    l.readChar()
  }
  return l.input[position:l.position]
}

//Schnauzer returns whether the passed byte is between 0 and 9 in isDigit
func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

//Shnauzer could care less about whitespace and uses Go RegEx to prove it!
func (l*Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.readChar()
  }
}

/*
Read Identifier of token ch bytes. For loops ch bytes in the string to see
if they are a letter and returns the array of every position
(every ch byte) but now as a string to be marked as an input string from the
Lexer data structure
*/

func (l *Lexer) readIdentifier() string {
  position := l.position
  for isLetter(l.ch) {
    l.readChar()
  }
  return l.input[position:l.position]
}

/*
Checks if the byte is a letter. Handles case sensitivity. Handles underscore.
Underscor so say var_foo would read.
*/

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch =='_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}
