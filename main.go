package main

import (
	"flag"
	"os"

	"github.com/jejikeh/gobf/bytecode"
)

type Run struct {
	*flag.FlagSet
	Input  *string
	Output *string
}

func NewRun() *Run {
	run := &Run{
		FlagSet: flag.NewFlagSet("build", flag.ExitOnError),
	}

	run.Input = run.String("i", "", "provide a input file")
	run.Output = run.String("o", "a.out", "output file")

	return run
}

func main() {
	runCmd := NewRun()

	switch os.Args[1] {
	case "run":
		runCmd.Parse(os.Args[2:])
		runCmd.compile()
	}
}

func (b *Run) compile() {
	p, err := bytecode.Generate(*b.Input)

	if err != nil {
		panic(err)
	}

	r := newRunner()
	r.run(p)
}
