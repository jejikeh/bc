package main

import "fmt"

type Operation struct {
	Kind    Token
	Operand int
}

type Program []Operation

type Compiler struct {
	Instructions  Program
	AddressBuffer []int
}

func newCompiler() *Compiler {
	return &Compiler{
		Instructions:  make(Program, 0),
		AddressBuffer: make([]int, 0),
	}
}

func (c *Compiler) compile(input string) (Program, error) {
	l := newLexer(input)
	v := l.getNext()

	lastOperation := Operation{
		Kind: End,
	}

	for v.Kind != End {
		op := Operation{
			Kind:    v.Kind,
			Operand: 1,
		}

		// [NOTE]: The instructions in Program a compressed. But this compression works
		// only for some tokens which are stackable.
		if v.Stackable && v.Kind == lastOperation.Kind {
			lastOperation.Operand++
			c.Instructions[len(c.Instructions)-1] = lastOperation
			v = l.getNext()

			continue
		}

		if v.Kind == JumpZero {
			address := len(c.Instructions)
			c.AddressBuffer = append(c.AddressBuffer, address)
		}

		if v.Kind == JumpNonZero {
			if len(c.AddressBuffer) == 0 {
				return nil, fmt.Errorf("stack underflow, JumpNonZero is reference to undefined stack value")
			}
		}

		lastOperation = op
		c.Instructions = append(c.Instructions, op)
		v = l.getNext()
	}

	return c.Instructions, nil
}

func (c *Compiler) trace() {
	for i, v := range c.Instructions {
		fmt.Printf("[%d]\t %c\t %d\n", i, v.Kind, v.Operand)
	}
}
