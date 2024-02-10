package bytecode

import (
	"os"
	"testing"
)

func TestLexCheckHardCodeSample(t *testing.T) {
	path := "../samples/hello.b"
	fileContent, err := os.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	tokens := lex(fileContent)

	if len(tokens) != len(helloKinds) {
		t.Errorf("Expected %d tokens but got %d", len(helloKinds), len(tokens))
	}

	for i, token := range tokens {
		if token.Kind != helloKinds[i] {
			t.Errorf("Expected '%c' but got '%c'", helloKinds[i], token)
		}
	}
}

func TestCheckUnknownChar(t *testing.T) {
	assertTokens := []Kind{
		Increment,
		Increment,
		Increment,
		MarkJump,
		Decrement,
		Decrement,
		Jump,
		End,
	}
	content := []byte("+++[--hello world]")

	tokens := lex(content)

	if len(tokens) != len(assertTokens) {
		t.Fatalf("Expected %d tokens but got %d", len(assertTokens), len(tokens))
	}

	for i, token := range tokens {
		if token.Kind != assertTokens[i] {
			t.Errorf("Expected %c but got %c", assertTokens[i], token)
		}
	}
}
