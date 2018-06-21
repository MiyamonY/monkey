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

	"github.com/MiyamonY/monkey/lexer"
	"github.com/MiyamonY/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
