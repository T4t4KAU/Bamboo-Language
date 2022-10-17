package command

import (
	"bamboo/evaluator"
	"bamboo/lexer"
	"bamboo/object"
	"bamboo/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, flag bool) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	if flag == true {
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
				PrintParserErrors(out, p.Errors())
				continue
			}

			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				output := evaluated.Inspect()
				if output == "NULL" {
					output = ""
				}
				io.WriteString(out, output)
				io.WriteString(out, "\n")
			}
		}
	} else {
		code := ""
		for scanner.Scan() {
			code += scanner.Text()
		}
		lex := lexer.New(code)
		p := parser.New(lex)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			PrintParserErrors(out, p.Errors())
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			output := evaluated.Inspect()
			if output == "NULL" {
				output = ""
			}
			io.WriteString(out, output)
			io.WriteString(out, "\n")
		}
	}
}

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser errors: ")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
