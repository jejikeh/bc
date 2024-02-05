package main

import "testing"

func TestReturnUnknowEmptyString(t *testing.T) {
	lexer := newLexer("")
	r := lexer.getNext()

	if r != Unknown {
		t.Error("empty string should return unknown operation kind!")
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

	if null != Unknown {
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
		if null != Unknown {
			t.Error("+- should return unknown operation kind!")
		}
	}
}
