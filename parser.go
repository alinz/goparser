package goparser

// AstNode defines an interface for asbtarct sytaxt tree
type AstNode interface {
	// Type returns the type of AstNode. This is a fast way to detect
	// the custom type of node so the value can be cast into the correct one
	Type() int
	// Value returns the value of current Node. It is a best practice to
	// use the type as check before performing cast to custom type.
	Value() interface{}
}

// Parser this is the main interface to build a parser
type Parser interface {
	// Parse gets the Tokens and returns the Ast presenation of
	// tokens. It would be an ideal if the implementation of Parser
	// is stateless so it can be used by multiple goroutines.
	Parse(Tokens) AstNode
}
