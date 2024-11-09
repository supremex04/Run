package repl

import (
	"bufio"
	"fmt"
	"io"
	"run/evaluator"
	"run/lexer"
	"run/parser"
)

const RUN_BANNER = `
      ___           ___           ___     
     /  /\         /__/\         /__/\    
    /  /::\        \  \:\        \  \:\   
   /  /:/\:\        \  \:\        \  \:\  
  /  /:/~/:/    ___  \  \:\   _____\__\:\ 
 /__/:/ /:/___ /__/\  \__\:\ /__/::::::::\
 \  \:\/:::::/ \  \:\ /  /:/ \  \:\~~\~~\/
  \  \::/~~~~   \  \:\  /:/   \  \:\  ~~~ 
   \  \:\        \  \:\/:/     \  \:\     
    \  \:\        \  \::/       \  \:\    
     \__\/         \__\/         \__\/    
`
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Printf(RUN_BANNER)
	fmt.Printf("\nType commands here:\n")

	for {
		fmt.Printf(PROMPT)

		// The scanner.Scan() method reads the next line from the input,
		//breaking the input into tokens based on newlines
		// (by default, it uses a newline \n as the delimiter).
		scanned := scanner.Scan()
		// The scanner.Scan() call returns true if there is another line to read,
		//or false if the input has ended (like if EOF is encountered).
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some errors!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
