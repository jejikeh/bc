package bytecode

import (
	"fmt"
	"os"
	"testing"
)

type IntermediateRepresentation struct {
	Token Token
	Op    int
}

func generate(path string) ([]IntermediateRepresentation, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	ir := make([]IntermediateRepresentation, 0)
	jumps := make([]int, 0)

	for _, r := range lex(fileContent) {
		switch r.Kind {
		case Increment:
			ir = appendOrIncrement(ir, r)
		case Decrement:
			ir = appendOrIncrement(ir, r)
		case Left:
			ir = appendOrIncrement(ir, r)
		case Right:
			ir = appendOrIncrement(ir, r)
		case Input:
			ir = appendOrIncrement(ir, r)
		case Output:
			ir = appendOrIncrement(ir, r)
		case MarkJump:
			jumps = append(jumps, len(ir))
			ir = append(ir, IntermediateRepresentation{Token: r, Op: 0})
		case Jump:
			if len(jumps) == 0 {
				return nil, fmt.Errorf("[%s:%d] Error: Jump without MarkJump", path, r.Position)
			}

			lastJump := jumps[len(jumps)-1]
			jumps = jumps[:len(jumps)-1]

			ir[lastJump].Op = len(ir)

			ir = append(ir, IntermediateRepresentation{Token: r, Op: lastJump})
		case End:
			ir = append(ir, IntermediateRepresentation{Token: r, Op: 0})
		}
	}

	return ir, nil
}

func appendOrIncrement(ir []IntermediateRepresentation, token Token) []IntermediateRepresentation {
	if len(ir) == 0 || ir[len(ir)-1].Token.Kind != token.Kind {
		ir = append(ir, IntermediateRepresentation{Token: token, Op: 1})
	} else {
		ir[len(ir)-1].Op++
	}

	return ir
}

func TestGenerateCheckHardCodeSample(t *testing.T) {
	path := "../samples/hello.b"
	b, err := generate(path)

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
