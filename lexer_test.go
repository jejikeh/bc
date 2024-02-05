package main

import "testing"

func TestReturnUnknowEmptyString(t *testing.T) {
	lexer := newLexer("")
	r := lexer.getNext()

	if r != End {
		t.Error("empty string should return End operation kind!")
	}
}

func TestOnlyIncrement(t *testing.T) {
	lexer := newLexer("+")
	r := lexer.getNext()

	if r != Increment {
		t.Error("+ should return incerement operation kind!")
	}
}

func TestOnlyDecrement(t *testing.T) {
	lexer := newLexer("-")
	r := lexer.getNext()

	if r != Decrement {
		t.Error("- should return decrement operation kind!")
	}
}

func TestIncrementAndDecrement(t *testing.T) {
	lexer := newLexer("+-")

	inc := lexer.getNext()
	dec := lexer.getNext()

	if inc != Increment {
		t.Error("+- should return increment operation kind!")
	}

	if dec != Decrement {
		t.Error("+- should return decrement operation kind!")
	}
}

func TestReturnUnknow(t *testing.T) {
	lexer := newLexer("+-h-[]")

	_ = lexer.getNext()
	_ = lexer.getNext()

	unk := lexer.getNext()

	_ = lexer.getNext()

	if unk != Decrement {
		t.Error("+-h-[] should skip 'h' and return decrement operation kind!")
	}
}

func TestReturnUnknowAtEnd(t *testing.T) {
	lexer := newLexer("+-")

	inc := lexer.getNext()
	dec := lexer.getNext()
	null := lexer.getNext()

	if inc != Increment {
		t.Error("+- should return increment operation kind!")
	}

	if dec != Decrement {
		t.Error("+- should return decrement operation kind!")
	}

	if null != End {
		t.Error("+- should return unknown operation kind!")
	}
}

func TestReturnUnknowAlways(t *testing.T) {
	lexer := newLexer("+-")

	inc := lexer.getNext()
	dec := lexer.getNext()

	if inc != Increment {
		t.Error("+- should return increment operation kind!")
	}

	if dec != Decrement {
		t.Error("+- should return decrement operation kind!")
	}

	for i := 0; i < 10; i++ {
		null := lexer.getNext()
		if null != End {
			t.Error("+- should return End operation kind!")
		}
	}
}

func TestShouldReturnAllTokens(t *testing.T) {
	helloWorld := "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+."

	lexer := newLexer(helloWorld)

	for i := 0; i < len(helloWorld)-1; i++ {
		if lexer.getNext() == End {
			t.Error("should return all tokens!")
		}
	}
}
