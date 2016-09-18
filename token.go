package goparser

// Token is an interface which defines all the types related to
// data out from lexer. this which consues by parse as input to
// constrcut the AST
type Token interface {
	// Value returns the string value of lex's token
	Value() string
	// Type returns the type of token.
	Type() int
}

// Tokens is an interface to abstratc the complexity of
// extracting each token. All the logic of custom lexer must be encapsoluated by
// this interface and pass to Parser object.
type Tokens interface {
	// Current returns the current token. it does not advanced the internal pointer
	Current() Token

	// Next returns the next token and advanced the internal pos by one
	Next() Token

	// PeekNth returns the nth token. it can be nil if
	// reaches end of tokens.
	PeekNth(int) Token
}

type defaultToken struct {
	val string
	typ int
}

func (s defaultToken) Value() string {
	return s.val
}

func (s defaultToken) Type() int {
	return s.typ
}

// NewToken creates a token based on given value and type
func NewToken(val string, typ int) Token {
	return &defaultToken{val: val, typ: typ}
}
