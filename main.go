package main

import (
	"flag"
	"os"
)

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
