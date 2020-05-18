package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Failure exit codes.
const (
	EXITCODE_FOPEN = 1
	EXITCODE_FSTAT = 2
	EXITCODE_FREAD = 3
	EXITCODE_LEX   = 4
	EXITCODE_PARSE = 5
)

// Abort with a message, an underlying error, and an exit code.
func die(msg string, err error, code int) {
	fmt.Fprintf(os.Stderr, "Error: %s: %s\n", msg, err)
	os.Exit(code)
}

// This function is the program's entry point.
func main() {
	rand.Seed(time.Now().UnixNano())

	mustCompileRegexes()

	// create a hard-coded env for now.
	globalBindings := make(Bindings)
	globalBindings["foo"] = LispInt(5)
	env := EnvStack{
		Bindings: globalBindings,
		Next:     nil,
	}

	var bytes []byte
	if len(os.Args) > 1 {
		fname := os.Args[len(os.Args)-1]
		bytes = mustLoadFile(fname)
	}

	tokens := mustLex(bytes)
	ast := mustParse(tokens)
	value := eval(ast, &env)
	_ = value
	fmt.Printf("%s\n", value.LispValueString())
}
