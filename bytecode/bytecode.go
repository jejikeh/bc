package bytecode

import (
	"fmt"
	"os"
)

type IntermediateRepresentation struct {
	Token Token
	Op    int
}

func Generate(path string) ([]IntermediateRepresentation, error) {
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
