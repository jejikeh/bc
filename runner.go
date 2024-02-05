package main

import "fmt"

type Runner struct {
	Run                Program
	Memory             []rune
	Head               int
	InstructionPointer int
}

func newRunner(program Program) Runner {
	return Runner{
		Run:    program,
		Memory: make([]rune, 2),
	}
}

func (r *Runner) exec() {
	for r.InstructionPointer < len(r.Run) {
		v := r.Run[r.InstructionPointer]
		switch v.Kind {
		case Increment:
			r.Memory[r.Head] += rune(v.Operand)
			r.InstructionPointer += 1
		case Decrement:
			r.Memory[r.Head] -= rune(v.Operand)
			r.InstructionPointer += 1
		case Left:
			if r.Head < r.Run[r.InstructionPointer].Operand {
				panic("Runtime Error. Memory underflow")
			}
			r.Head -= v.Operand
			r.InstructionPointer += 1
		case Right:
			r.Head += v.Operand

			for r.Head >= len(r.Memory) {
				r.Memory = append(r.Memory, 1)
			}

			r.InstructionPointer += 1
		case Input:
		case Output:
			fmt.Printf("%c", r.Memory[r.Head])
			r.InstructionPointer += 1
		case JumpZero:
			if r.Memory[r.Head] == 0 {
				r.InstructionPointer = r.Run[r.InstructionPointer].Operand
			} else {
				r.InstructionPointer += 1
			}
		case JumpNonZero:
			if r.Memory[r.Head] != 0 {
				r.InstructionPointer = r.Run[r.InstructionPointer].Operand
			} else {
				r.InstructionPointer += 1
			}
		}
	}
}
