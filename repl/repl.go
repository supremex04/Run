package repl

import (
	"bufio"
	"fmt"
	"io"
	"run/lexer"
	"run/token"
)
const PROMPT = ">> "


func Start(in io.Reader /*, out io.Writer*/) {
	scanner := bufio.NewScanner(in)

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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
			