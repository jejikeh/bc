package main

import "testing"

func Test_initCompiler(t *testing.T) {
	compiler := newCompiler()

	err := compiler.compile("hello, word")

	if err != nil {
		t.Error("Err is not null")
	}
}
