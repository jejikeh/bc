package main

import (
	"flag"
	"os"
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
	build.Trace = build.Bool("t", false, "trace the program")

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
	content, err := os.ReadFile(*b.Input)

	if err != nil {
		panic(err)
	}

	c := newCompiler()
	p, err := c.compile(string(content))

	if err != nil {
		panic(err)
	}

	if *b.Trace {
		c.trace()
	}

	r := newRunner(p)
	r.exec()
}
