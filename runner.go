package main

import (
	"fmt"

	"github.com/jejikeh/gobf/bytecode"
)

type Runner struct {
	Memory             []rune
	Head               int
	InstructionPointer int
}

func newRunner() Runner {
	return Runner{
		Memory: make([]rune, 2),
	}
}

func (r *Runner) run(program []bytecode.IntermediateRepresentation) {
	for r.InstructionPointer < len(program) {
		v := program[r.InstructionPointer]
		switch v.Token.Kind {
		case bytecode.Increment:
			r.Memory[r.Head] += rune(v.Op)
			r.InstructionPointer += 1
		case bytecode.Decrement:
			r.Memory[r.Head] -= rune(v.Op)
			r.InstructionPointer += 1
		case bytecode.Left:
			if r.Head < program[r.InstructionPointer].Op {
				panic("Runtime Error. Memory underflow")
			}
			r.Head -= v.Op
			r.InstructionPointer += 1
		case bytecode.Right:
			r.Head += v.Op

			for r.Head >= len(r.Memory) {
				r.Memory = append(r.Memory, 1)
			}

			r.InstructionPointer += 1
		case bytecode.Input:
		case bytecode.Output:
			fmt.Printf("%c", r.Memory[r.Head])
			r.InstructionPointer += 1
		case bytecode.MarkJump:
			if r.Memory[r.Head] == 0 {
				r.InstructionPointer = program[r.InstructionPointer].Op
			} else {
				r.InstructionPointer += 1
			}
		case bytecode.Jump:
			if r.Memory[r.Head] != 0 {
				r.InstructionPointer = program[r.InstructionPointer].Op
			} else {
				r.InstructionPointer += 1
			}
		}
	}
}
