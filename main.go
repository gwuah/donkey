package main

import (
	"fmt"

	lexer "github.com/gwuah/donkey/lexer"
)

func print(value interface{}) {
	fmt.Println(value)
}

func main() {
	input := `=+(){},;`
	l := lexer.New(input)

	l.NextToken()

}
