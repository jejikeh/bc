package main

import (
	"flag"
	"os"

	"github.com/jejikeh/gobf/bytecode"
)

type Build struct {
	*flag.FlagSet
	Input  *string
	Output *string
	Trace  *bool
}

func NewBuild() *Build {
	build := &Build{
		FlagSet: flag.NewFlagSet("build", flag.ExitOnError),
	}

	build.Input = build.String("i", "", "provide a input file")
	build.Output = build.String("o", "a.out", "output file")

	return build
}

func main() {
	buildCmd := NewBuild()

	switch os.Args[1] {
	case "build":
		buildCmd.Parse(os.Args[2:])
		buildCmd.compile()
	}
}

func (b *Build) compile() {
	p, err := bytecode.Generate(*b.Input)

	if err != nil {
		panic(err)
	}

	r := newRunner()
	r.run(p)
}
