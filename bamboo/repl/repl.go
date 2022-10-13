package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

const PROMPT = ">> "
const message = " ____                  _                 \n| __ )  __ _ _ __ ___ | |__   ___   ___  \n|  _ \\ / _` | '_ ` _ \\| '_ \\ / _ \\ / _ \\ \n| |_) | (_| | | | | | | |_) | (_) | (_) |\n|____/ \\__,_|_| |_| |_|_.__/ \\___/ \\___/ \n\n"

func Start(in io.Reader, out io.Writer) {
	fmt.Printf(message)
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lex := lexer.New(line)
		p := parser.New(lex)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "ran into some monkey\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
