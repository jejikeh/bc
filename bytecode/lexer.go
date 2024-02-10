package bytecode

type Token struct {
	Kind     Kind
	Position int
}

type Kind rune

const (
	Increment Kind = '+'
	Decrement Kind = '-'
	Left      Kind = '<'
	Right     Kind = '>'
	Input     Kind = ','
	Output    Kind = '.'
	MarkJump  Kind = '['
	Jump      Kind = ']'
	End       Kind = '∅'
)

func lex(content []byte) []Token {
	Kinds := make([]Token, 0)

	for i, r := range content {
		switch r {
		case byte(Increment):
			Kinds = append(Kinds, Token{Increment, i})
		case byte(Decrement):
			Kinds = append(Kinds, Token{Decrement, i})
		case byte(Left):
			Kinds = append(Kinds, Token{Left, i})
		case byte(Right):
			Kinds = append(Kinds, Token{Right, i})
		case byte(Input):
			Kinds = append(Kinds, Token{Input, i})
		case byte(Output):
			Kinds = append(Kinds, Token{Output, i})
		case byte(MarkJump):
			Kinds = append(Kinds, Token{MarkJump, i})
		case byte(Jump):
			Kinds = append(Kinds, Token{Jump, i})
		}
	}

	Kinds = append(Kinds, Token{End, len(content)})

	return Kinds
}
