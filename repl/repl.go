///
// File:  repl.go
// Author: ymiyamoto
//
// Created on Thu Jun 21 09:48:13 2018
//
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MiyamonY/monkey/evaluator"
	"github.com/MiyamonY/monkey/lexer"
	"github.com/MiyamonY/monkey/object"
	"github.com/MiyamonY/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if evaluated := evaluator.Eval(program, env); evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
