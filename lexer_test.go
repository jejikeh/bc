package main

import "testing"

func TestReturnUnknowEmptyString(t *testing.T) {
	lexer := newLexer("")
	r := lexer.getNext()

	if r.Kind != End {
		t.Error("empty string should return End operation kind!")
	}
}

func TestOnlyIncrement(t *testing.T) {
	lexer := newLexer("+")
	r := lexer.getNext()

	if r.Kind != Increment {
		t.Error("+ should return incerement operation kind!")
	}
}

func TestOnlyDecrement(t *testing.T) {
	lexer := newLexer("-")
	r := lexer.getNext()

	if r.Kind != Decrement {
		t.Error("- should return decrement operation kind!")
	}
}

func TestIncrementAndDecrement(t *testing.T) {
	lexer := newLexer("+-")

	inc := lexer.getNext()
	dec := lexer.getNext()

	if inc.Kind != Increment {
		t.Error("+- should return increment operation kind!")
	}

	if dec.Kind != Decrement {
		t.Error("+- should return decrement operation kind!")
	}
}

func TestReturnUnknown(t *testing.T) {
	lexer := newLexer("+-h-[]")

	_ = lexer.getNext()
	_ = lexer.getNext()

	unk := lexer.getNext()

	_ = lexer.getNext()

	if unk.Kind != Decrement {
		t.Error("+-h-[] should skip 'h' and return decrement operation kind!")
	}
}

func TestReturnUnknownAtEnd(t *testing.T) {
	lexer := newLexer("+-")

	inc := lexer.getNext()
	dec := lexer.getNext()
	null := lexer.getNext()

	if inc.Kind != Increment {
		t.Error("+- should return increment operation kind!")
	}

	if dec.Kind != Decrement {
		t.Error("+- should return decrement operation kind!")
	}

	if null.Kind != End {
		t.Error("+- should return unknown operation kind!")
	}
}

func TestReturnUnknownAlways(t *testing.T) {
	lexer := newLexer("+-")

	inc := lexer.getNext()
	dec := lexer.getNext()

	if inc.Kind != Increment {
		t.Error("+- should return increment operation kind!")
	}

	if dec.Kind != Decrement {
		t.Error("+- should return decrement operation kind!")
	}

	for i := 0; i < 10; i++ {
		null := lexer.getNext()
		if null.Kind != End {
			t.Error("+- should return End operation kind!")
		}
	}
}

func TestShouldReturnAllTokens(t *testing.T) {
	helloWorld := "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+."

	lexer := newLexer(helloWorld)

	for i := 0; i < len(helloWorld)-1; i++ {
		if lexer.getNext().Kind == End {
			t.Error("should return all tokens!")
		}
	}
}
