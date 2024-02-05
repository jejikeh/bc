package main

type OperationKind struct {
	Kind      Token
	Stackable bool
}

type Token rune

const (
	Increment   Token = '+'
	Decrement   Token = '-'
	Left        Token = '<'
	Right       Token = '>'
	Input       Token = ','
	Output      Token = '.'
	JumpZero    Token = '['
	JumpNonZero Token = ']'
	End         Token = '∅'
)

func newOperationKind(kind Token, stackable bool) OperationKind {
	return OperationKind{
		Kind:      kind,
		Stackable: stackable,
	}
}

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
	for l.Position < len(l.Content) {
		r := l.Content[l.Position]
		l.Position++

		switch r {
		case byte(Increment):
			return newOperationKind(Increment, true)
		case byte(Decrement):
			return newOperationKind(Decrement, true)
		case byte(Left):
			return newOperationKind(Left, true)
		case byte(Right):
			return newOperationKind(Right, true)
		case byte(Input):
			return newOperationKind(Input, true)
		case byte(Output):
			return newOperationKind(Output, true)
		case byte(JumpZero):
			return newOperationKind(JumpZero, false)
		case byte(JumpNonZero):
			return newOperationKind(JumpNonZero, false)
		default:
			return l.getNext()
		}
	}

	return newOperationKind(End, false)
}
