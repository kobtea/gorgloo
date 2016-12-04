package gorgloo

// Token represents a lexical token
type Token int

const (
	// ILLEGAL represents undefined token
	ILLEGAL Token = iota
	// EOF represents end of files
	EOF
	// WS represents white spaces
	WS

	// ASTERISK represents `*`s
	ASTERISK
	// STATE represents a task state
	STATE
	// PRIORITY represents a task priority
	PRIORITY
	// HEADLINE represents a task summary
	HEADLINE
	// TAG represents a task category
	TAG
)
