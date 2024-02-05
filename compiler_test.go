package main

import (
	"os"
	"testing"
)

func TestInitCompiler(t *testing.T) {
	compiler := newCompiler()

	_, err := compiler.compile("hello, word")

	if err != nil {
		t.Error("Err is not null")
	}
}

func TestCompressSameOperations(t *testing.T) {
	compiler := newCompiler()

	p, err := compiler.compile("++++++")

	if err != nil {
		t.Error("Err is not null")
	}

	if len(p) != 1 {
		t.Error("length should be 1")
	}

	if p[0].Kind != Increment {
		t.Error("Kind should be Increment")
	}

	if p[0].Operand != 6 {
		t.Errorf("Operand should be 5, got %d", p[0].Operand)
	}
}

func TestCompressSameOperationsMore(t *testing.T) {
	compiler := newCompiler()

	p, err := compiler.compile("++++++----")

	if err != nil {
		t.Error("Err is not null")
	}

	if len(p) != 2 {
		t.Error("length should be 2")
	}

	if p[0].Kind != Increment {
		t.Error("Kind should be Increment")
	}

	if p[0].Operand != 6 {
		t.Errorf("Operand should be 6, got %d", p[0].Operand)
	}

	if p[1].Kind != Decrement {
		t.Error("Kind should be Decrement")
	}

	if p[1].Operand != 4 {
		t.Errorf("Operand should be 4, got %d", p[1].Operand)
	}
}

func TestDoNotCompressDifferentOperations(t *testing.T) {
	compiler := newCompiler()
	p, err := compiler.compile("+-[]")

	if err != nil {
		t.Error("Err is not null")
	}

	if len(p) != 4 {
		t.Error("length should be 4")
	}

	if p[0].Kind != Increment {
		t.Error("Kind should be Increment")
	}

	if p[0].Operand != 1 {
		t.Errorf("Operand should be 1, got %d", p[0].Operand)
	}

	if p[1].Kind != Decrement {
		t.Error("Kind should be Decrement")
	}

	if p[1].Operand != 1 {
		t.Errorf("Operand should be 4, got %d", p[1].Operand)
	}

	if p[2].Kind != JumpZero {
		t.Error("Kind should be JumpZero")
	}

	if p[3].Kind != JumpNonZero {
		t.Error("Kind should be JumpNonZero")
	}
}

func TestHelloWorld(t *testing.T) {
	compiler := newCompiler()
	samples, err := os.ReadDir("samples")

	if err != nil {
		t.Error("Error while opening samples directory")
	}

	for sample := range samples {
		sampleContent, err := os.ReadFile("samples/" + samples[sample].Name())

		if err != nil {
			t.Error("Error while opening", samples[sample].Name())
		}

		_, err = compiler.compile(string(sampleContent))

		if err != nil {
			t.Error("Error while compiling", samples[sample].Name())
		}
	}
}
