package main

import (
	"flag"
	"os"
)

type OperationKind int

const (
	Increment OperationKind = iota
	Decrement
	Left
	Right
	Input
	Output
	JumpZero
	JumpNonZero
	Unknown
)

type Operation struct {
	Kind    OperationKind
	Operand int
}

type Program []Operation

func main() {
	inputFile := flag.String("i", "", "provide a input file")
	outputFile := flag.String("o", "a.out", "output file")

	flag.Parse()

	compile(*inputFile, *outputFile)
}

func compile(input string, output string) {
	content, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	c := newCompiler()

	c.compile(string(content))
}
