package main

type Lexer struct {
	Content  string
	Position int
}

func newLexer(content string) *Lexer {
	return &Lexer{
		Content: content,
	}
}

func (l *Lexer) getNext() OperationKind {
	if l.Position < len(l.Content) {
		r := l.Content[l.Position]
		l.Position++
		return l.getKind(r)
	}

	return Unknown
}

func (l *Lexer) getKind(r byte) OperationKind {
	switch r {
	case '+':
		return Increment
	case '-':
		return Decrement
	}

	return Unknown
}
