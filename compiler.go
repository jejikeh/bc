package main

import "fmt"

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

var lastOperation *Operation = nil

func (c *Compiler) compile(input string) (Program, error) {
	l := newLexer(input)
	v := l.getNext()
	for v != End {
		op := Operation{
			Kind:    v,
			Operand: 1,
		}

		if lastOperation != nil && (v == Increment || v == Decrement || v == Left || v == Right || v == Input || v == Output) {
			if lastOperation.Kind == v {
				lastOperation.Operand++
				c.Instructions[len(c.Instructions)-1] = *lastOperation
				v = l.getNext()
				continue
			}
		}

		lastOperation = &op
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
