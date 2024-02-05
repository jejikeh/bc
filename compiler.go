package main

type Compiler struct {
	Instructions Program
}

func newCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) compile(input string) error {
	for _, v := range input {
		switch v {
		case '+':
			op := Operation{
				Kind:    Increment,
				Operand: 1,
			}
			c.Instructions = append(c.Instructions, op)
		case '-':
			op := Operation{
				Kind:    Decrement,
				Operand: 1,
			}

			c.Instructions = append(c.Instructions, op)
		}
	}

	return nil
}
