package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gwuah/donkey/lexer"
	"github.com/gwuah/donkey/token"
)

const PROMPT = "\n🦙 >> "

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

		for tok := l.NextToken(); tok.Type != token.EOF; {
			fmt.Printf("%+v\n", tok)
			tok = l.NextToken()
		}
	}

}
