package parser

import (
	"github.com/gwuah/donkey/lexer"
	"github.com/gwuah/donkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}
