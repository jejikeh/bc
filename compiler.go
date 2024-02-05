package main

type Operation struct {
	Kind    OperationKind
	Operand int
}

type Program []Operation

type Compiler struct {
	Instructions Program
}

func newCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) compile(input string) (Program, error) {
	l := newLexer(input)
	v := l.getNext()
	for v != End {
		op := Operation{
			Kind: v,
		}

		if v == Increment || v == Decrement || v == Left || v == Right || v == Input || v == Output {
			if len(c.Instructions) == 0 {
				op.Operand = 1
			} else {
				last := c.Instructions[len(c.Instructions)-1]
				if last.Kind == v {
					last.Operand++
					c.Instructions[len(c.Instructions)-1] = last
					v = l.getNext()
					continue
				} else {
					op.Operand = 1
				}
			}
		}

		c.Instructions = append(c.Instructions, op)
		v = l.getNext()
	}

	return c.Instructions, nil
}
