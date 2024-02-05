package main

type OperationKind rune

const (
	Increment   OperationKind = '+'
	Decrement   OperationKind = '-'
	Left        OperationKind = '<'
	Right       OperationKind = '>'
	Input       OperationKind = ','
	Output      OperationKind = '.'
	JumpZero    OperationKind = '['
	JumpNonZero OperationKind = ']'
	End         OperationKind = '∅'
)

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
			return Increment
		case byte(Decrement):
			return Decrement
		case byte(Left):
			return Left
		case byte(Right):
			return Right
		case byte(Input):
			return Input
		case byte(Output):
			return Output
		case byte(JumpZero):
			return JumpZero
		case byte(JumpNonZero):
			return JumpNonZero
		default:
			return l.getNext()
		}
	}

	return End
}
