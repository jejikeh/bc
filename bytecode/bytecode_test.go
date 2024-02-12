package bytecode

import (
	"testing"
)

func TestGenerateCheckHardCodeSample(t *testing.T) {
	path := "../samples/hello.b"
	b, err := Generate(path)

	if err != nil {
		t.Error(err)
	}

	if len(helloIntermediateRepresentation) != len(b) {
		t.Errorf("Expected %d tokens but got %d", len(helloIntermediateRepresentation), len(b))
	}

	for i, ir := range b {
		if ir.Token.Kind != helloIntermediateRepresentation[i].Token.Kind {
			t.Errorf("Expected Kind '%c' but got '%c' at %d", helloIntermediateRepresentation[i].Token.Kind, ir.Token.Kind, ir.Token.Position)
		}

		if ir.Op != helloIntermediateRepresentation[i].Op {
			t.Errorf("Expected Op %d but got %d at %d", helloIntermediateRepresentation[i].Op, ir.Op, ir.Token.Position)
		}

		if ir.Token.Position != helloIntermediateRepresentation[i].Token.Position {
			t.Errorf("Expected Position %d but got %d at %d", helloIntermediateRepresentation[i].Token.Position, ir.Token.Position, ir.Token.Position)
		}
	}
}
